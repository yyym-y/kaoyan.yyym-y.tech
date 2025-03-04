package admin

import (
	"back/biz/config"
	"back/biz/middle"
	"back/biz/model"
	reqmodel "back/biz/reqModel"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"
)

type VedioConf struct {
	vedioBaseUrl string
}

var Vconf VedioConf

func init() {
	Vconf = VedioConf{}
	config.Add(&Vconf)
}

func (cof *VedioConf) Update() {
	cof.vedioBaseUrl = config.Conf.VedioPath
}

func removeFilesInDir(dir string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			err = os.Remove(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func DivVedio(VedioId int, Num int, oriFileName string) (string, error) {
	path := Vconf.vedioBaseUrl + strconv.Itoa(VedioId) + "/" + strconv.Itoa(Num) + "/"
	temPath := TSconf.temSaveBaseUrl + oriFileName
	fmt.Println(path)
	fmt.Println(temPath)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	removeFilesInDir(path)

	// 将其切分为 m3u8 文件
	finalName := middle.NewUUID() + ".m3u8"
	commend := "ffmpeg -i " + temPath +
		" -c copy -start_number 0 -f hls -hls_time 10 -hls_list_size 0 " +
		path + finalName
	fmt.Println(commend)
	cmd := exec.Command("cmd", "/C", commend)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		return "", err
	}
	return finalName, nil
}

func EditVedioHandler(ctx context.Context, c *app.RequestContext) {
	editVal := reqmodel.EditVedio{}
	if err := c.BindJSON(&editVal); err != nil {
		c.JSON(400, middle.FailWithMsg("bind wrong"))
	}
	fmt.Println(editVal)

	var wg sync.WaitGroup

	oriVedioInfo := model.Vedio{}
	model.DB.Find(&model.Vedio{}).Where("vedio_id = ?", editVal.VedioId).First(&oriVedioInfo)
	newVedioInfo := model.Vedio{
		VedioName:   editVal.VedioName,
		VedioId:     editVal.VedioId,
		VedioCover:  editVal.VedioCover,
		Description: editVal.Description,
		Visable:     editVal.Visable,
		TypeId:      editVal.TypeId,
	}

	// fmt.Println(newVedioInfo)

	// 处理封面修改的问题
	if oriVedioInfo.VedioCover != editVal.VedioCover {
		coverName := middle.NewUUID() + ".png"
		info, err := middle.ReadFile(editVal.VedioCover)
		if err != nil {
			c.JSON(400, middle.FailWithMsg("cannot read cover in cache"))
		}
		newVedioInfo.VedioCover = coverName
		wg.Add(1)
		go func() {
			defer wg.Done()
			writeImg(coverName, info)
		}()
	}
	// 处理视频被修改的问题
	oriEpiso := make([]model.PlayInfo, 0)
	model.DB.Where("vedio_id = ?", editVal.VedioId).Find(&oriEpiso)
	fmt.Println(editVal.VedioId)
	sort.Slice(oriEpiso, func(i, j int) bool {
		return oriEpiso[i].Num < oriEpiso[j].Num
	})
	sort.Slice(editVal.EpisodeInfos, func(i, j int) bool {
		return editVal.EpisodeInfos[i].Num < editVal.EpisodeInfos[j].Num
	})
	changeList := make([]model.PlayInfo, 0)
	changeNameList := make([]model.PlayInfo, 0)
	delList := make([]model.PlayInfo, 0)
	addList := make([]model.PlayInfo, 0)
	pos := 0
	for ; pos < len(editVal.EpisodeInfos); pos++ {
		if pos >= len(oriEpiso) {
			for ; pos < len(editVal.EpisodeInfos); pos++ {
				addList = append(addList, editVal.EpisodeInfos[pos])
			}
			break
		}
		if editVal.EpisodeInfos[pos] == oriEpiso[pos] {
			continue
		}
		if editVal.EpisodeInfos[pos].Url == oriEpiso[pos].Url {
			changeNameList = append(changeNameList, editVal.EpisodeInfos[pos])
			continue
		}
		changeList = append(changeList, editVal.EpisodeInfos[pos])
	}
	for ; pos < len(oriEpiso); pos++ {
		delList = append(delList, oriEpiso[pos])
	}

	fmt.Println(changeList)
	fmt.Println(changeNameList)
	fmt.Println(delList)
	fmt.Println(addList)

	addEpi := make(chan model.PlayInfo, 100)
	changeEpi := make(chan model.PlayInfo, 100)

	// 处理视频源给改变的
	for _, v := range changeList {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nname, _ := DivVedio(v.VedioId, v.Num, v.Url)
			v.Url = nname
			changeEpi <- v
		}()
	}
	// 处理只有其他信息（不包括视频）的信息
	for _, v := range changeNameList {
		wg.Add(1)
		go func() {
			defer wg.Done()
			changeEpi <- v
		}()
	}
	// 处理增加的信息
	for _, v := range addList {
		wg.Add(1)
		go func(v model.PlayInfo) {
			defer wg.Done()
			nname, _ := DivVedio(v.VedioId, v.Num, v.Url)
			v.Url = nname
			addEpi <- v
		}(v)
	}
	wg.Wait()
	close(addEpi)
	close(changeEpi)

	err := model.DB.Transaction(func(tx *gorm.DB) error {
		fmt.Println("into transaction")
		err := tx.Model(&model.Vedio{}).Where("vedio_id = ?", newVedioInfo.VedioId).Save(&newVedioInfo)
		if err.Error != nil {
			fmt.Println(err.Error)
			return err.Error
		}
		for v := range addEpi {
			err := tx.Create(&v)
			if err.Error != nil {
				fmt.Println(err.Error)
				return err.Error
			}
		}
		for v := range changeEpi {
			err := tx.Model(&model.PlayInfo{}).Where("vedio_id = ?", v.VedioId).Where("Num = ?", v.Num).Save(&v)
			if err.Error != nil {
				fmt.Println(err.Error)
				return err.Error
			}
		}
		return nil
	})
	if err != nil {
		c.JSON(500, middle.FailWithMsg("cannot write datebase"))
	}
	c.JSON(http.StatusOK, middle.SuccessWithNil())
}
