package vedio

import (
	"back/biz/config"
	"back/biz/middle"
	"context"
	"io"
	"net/http"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
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

func VedioPlayHandler(ctx context.Context, c *app.RequestContext) {
	vedioId := c.Param("VedioId")
	episo := c.Param("Episo")
	filename := c.Param("filename")
	path := Vconf.vedioBaseUrl + vedioId + "/" + episo + "/" + filename
	file, err := os.Open(path)
	if err != nil {
		c.JSON(http.StatusNotFound, middle.FailWithMsg("Vedio No Found"))
		return
	}
	defer file.Close()

	c.Header("Content-Type", "application/vnd.apple.mpegurl")
	c.Status(http.StatusOK) // 设置响应状态为 200 OK

	// 将文件内容写入响应
	if _, err := io.Copy(c, file); err != nil {
		c.JSON(http.StatusInternalServerError, middle.FailWithMsg("Failed to send Vedio"))
		return
	}
}
