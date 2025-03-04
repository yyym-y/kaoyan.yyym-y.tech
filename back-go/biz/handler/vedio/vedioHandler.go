package vedio

import (
	"back/biz/middle"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
)

func VedioHandler(ctx context.Context, c *app.RequestContext) {
	vedioId := c.Param("VedioId")
	episo := c.Param("Episo")
	filename := c.Param("filename")
	path := Vconf.vedioBaseUrl + vedioId + "/" + episo + "/" + filename
	file, err := os.Open(path)
	fmt.Println(path)
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
