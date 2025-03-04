package admin

import (
	"back/biz/config"
	"back/biz/middle"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
)

type LogConf struct {
	logBaseUrl string
}

var Lconf LogConf

func init() {
	Lconf = LogConf{}
	config.Add(&Lconf)
}

func (cof *LogConf) Update() {
	cof.logBaseUrl = config.Conf.LogPath
}

func GetDevelopContentHandler(ctx context.Context, c *app.RequestContext) {
	fullPath := filepath.Join(Lconf.logBaseUrl, "developLog.md")
	fmt.Println(fullPath)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		c.JSON(400, middle.FailWithMsg("读取文件失败"))
		return
	}
	fmt.Println(string(content))
	c.JSON(200, middle.SuccessWithDate(string(content)))
}

type LogSaveReq struct {
	Content string `json:"content"`
}

func SaveDevelopContentHandler(ctx context.Context, c *app.RequestContext) {
	req := LogSaveReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, middle.FailWithMsg("Error bind"))
		return
	}
	fullPath := filepath.Join(Lconf.logBaseUrl, "developLog.md")
	fmt.Println(fullPath)

	file, err := os.Create(fullPath) // 使用 os.Create 创建文件，如果文件已存在则会清空内容
	if err != nil {
		c.JSON(500, middle.FailWithMsg("Error creating file: "+err.Error()))
		return
	}
	defer file.Close() // 确保在函数结束时关闭文件
	// 写入内容到文件
	if _, err := file.WriteString(req.Content); err != nil {
		c.JSON(500, middle.FailWithMsg("Error writing to file: "+err.Error()))
		return
	}
	c.JSON(200, middle.SuccessWithNil())
}
