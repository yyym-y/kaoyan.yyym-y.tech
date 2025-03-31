package orc

import (
	"back/biz/config"
	"back/biz/middle"
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

type LatexConfig struct {
	LetexUAT      string
	LetexAppId    string
	LetexSecret   string
	LatexLightApi string
	LatexNomalApi string
	TemSavePath   string
}

var Lconf LatexConfig

func init() {
	Lconf = LatexConfig{}
	config.Add(&Lconf)
}

func (cof *LatexConfig) Update() {
	Lconf.LetexUAT = config.Conf.LetexUAT
	Lconf.LetexAppId = config.Conf.LetexAppId
	Lconf.LetexSecret = config.Conf.LetexSecret
	Lconf.LatexLightApi = config.Conf.LatexLightApi
	Lconf.LatexNomalApi = config.Conf.LatexNomalApi
	Lconf.TemSavePath = config.Conf.TemSavePath
}

func UploadImgHandler(ctx context.Context, c *app.RequestContext) {
	// 暂时不支持普通用户使用
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, middle.FailWithMsg("Error getting file"))
		return
	}
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, middle.FailWithMsg("Error open file"))
		return
	}
	defer src.Close()

	saveFileName := middle.NewUUID() + filepath.Ext(file.Filename)
	filePath := filepath.Join(Lconf.TemSavePath, saveFileName)

	// 写入文件
	dst, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating file"})
		return
	}
	defer dst.Close()

	// 将上传的文件内容复制到目标文件
	if _, err := io.Copy(dst, src); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error saving file"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, middle.SuccessWithDate(map[string]string{"fileName": saveFileName}))
}

func getReqData(reqData map[string]interface{}, appid, secret string) (http.Header, map[string]interface{}) {
	header := make(http.Header)
	header.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	header.Set("random-str", middle.NewUUID()[:14]+"Aa")
	header.Set("app-id", appid)

	preSignString := ""
	sortedKeys := make([]string, 0, len(reqData)+len(header))
	for k := range reqData {
		sortedKeys = append(sortedKeys, k)
	}
	for k := range header {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {
		if preSignString != "" {
			preSignString += "&"
		}
		lowerKey := strings.ToLower(key)
		if value, ok := header[key]; ok {
			preSignString += lowerKey + "=" + value[0]
		} else {
			preSignString += lowerKey + "=" + fmt.Sprintf("%v", reqData[key])
		}
	}

	preSignString += "&secret=" + secret
	fmt.Println(preSignString)
	hash := md5.Sum([]byte(preSignString))
	header.Set("sign", fmt.Sprintf("%x", hash))

	return header, reqData
}

func LatexOrcController(ctx context.Context, c *app.RequestContext) {
	type LatexOrcReq struct {
		FileName string `json:"fileName"`
		Model    int    `json:"model"`
	}
	req := LatexOrcReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, middle.FailWithMsg("Error JSON bind"))
		return
	}
	fmt.Println(req)

	// 开始调用 API
	data := make(map[string]interface{}) // 请求参数数据（非文件型参数），视情况填入

	header, data := getReqData(data, Lconf.LetexAppId, Lconf.LetexSecret)
	fmt.Println(header)
	fmt.Println(data)
	modelApi := Lconf.LatexNomalApi
	if req.Model == 1 {
		modelApi = Lconf.LatexLightApi
	}
	imgFilePath := filepath.Join(Lconf.TemSavePath, req.FileName)
	fmt.Println(imgFilePath)
	imgFileReader, err := os.Open(imgFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer imgFileReader.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", req.FileName)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}
	if _, err := io.Copy(part, imgFileReader); err != nil {
		fmt.Println("Error copying file to form:", err)
		return
	}
	// data 参数
	if recMode, ok := data["rec_mode"]; ok {
		writer.WriteField("rec_mode", fmt.Sprintf("%v", recMode))
	}

	writer.Close()

	// 创建一个新的 HTTP POST 请求
	apiReq, err := http.NewRequest("POST", modelApi, &buf)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	apiReq.Header.Set("Content-Type", writer.FormDataContentType())
	for key, values := range header {
		for _, value := range values {
			apiReq.Header.Add(key, value)
		}
	}
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(apiReq)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	// 打印响应
	fmt.Println(responseBody)
	fmt.Println("----------------")
	c.JSON(http.StatusOK, middle.SuccessWithDate(responseBody))
}
