package admin

import (
	"back/biz/config"
	"back/biz/middle"
	"back/biz/model"
	reqmodel "back/biz/reqModel"
	"bytes"
	"context"
	"io"
	"net/http"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
)

type ImgConf struct {
	imageBaseUrl string
}

var Iconf ImgConf

func init() {
	Iconf = ImgConf{}
	config.Add(&Iconf)
}

func (cof *ImgConf) Update() {
	cof.imageBaseUrl = config.Conf.ImgPath
}

func writeImg(fileName string, info []byte) error {
	path := Iconf.imageBaseUrl + fileName
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, bytes.NewReader(info)); err != nil {
		return err
	}
	return nil
}

func CreateVedioHandler(ctx context.Context, c *app.RequestContext) {
	req := reqmodel.CreateVedioReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, middle.FailWithMsg("bind wrong"))
		return
	}
	imgs, err := middle.ReadFile(req.VedioCover)
	if err != nil {
		c.JSON(400, middle.FailWithMsg("Cache find miss"))
		return
	}
	fileName := middle.NewUUID() + ".png"

	errChan := make(chan error)
	go func() {
		errChan <- writeImg(fileName, imgs)
	}()

	go func() {
		err := <-errChan
		if err != nil {
			c.JSON(500, middle.FailWithMsg("Failed to write image: "+err.Error()))
		} else {
			result := model.DB.Create(model.Vedio{
				VedioName:   req.VedioName,
				VedioCover:  fileName,
				Visable:     req.Visable,
				Description: req.Description,
				TypeId:      req.TypeId,
			})
			if result.Error != nil {
				c.JSON(500, middle.FailWithMsg("Failed to write datwbase"))
			}
			c.JSON(http.StatusOK, middle.SuccessWithNil())
		}
	}()
}
