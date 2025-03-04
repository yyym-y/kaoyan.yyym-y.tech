package index

import (
	"back/biz/config"
	"back/biz/middle"
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

func ImageHandler(ctx context.Context, c *app.RequestContext) {
	filename := c.Param("filename")
	file, err := os.Open(Iconf.imageBaseUrl + filename) // 假设图片存放在 images 文件夹中
	if err != nil {
		c.JSON(http.StatusNotFound, middle.FailWithMsg("Image No Found"))
		return
	}
	defer file.Close()

	c.Header("Content-Type", "image/jpeg") // 根据实际图片类型设置 Content-Type
	c.Status(http.StatusOK)                // 设置响应状态为 200 OK

	// 将文件内容写入响应
	if _, err := io.Copy(c, file); err != nil {
		c.JSON(http.StatusInternalServerError, middle.FailWithMsg("Failed to send image"))
		return
	}
}
