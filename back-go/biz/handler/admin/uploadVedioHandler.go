package admin

import (
	"back/biz/config"
	"back/biz/middle"
	"context"
	"io"
	"net/http"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
)

type TemSaveConf struct {
	temSaveBaseUrl string
}

var TSconf TemSaveConf

func init() {
	TSconf = TemSaveConf{}
	config.Add(&TSconf)
}

func (cof *TemSaveConf) Update() {
	cof.temSaveBaseUrl = config.Conf.TemSavePath
}

func UploadVedioHandler(ctx context.Context, c *app.RequestContext) {
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
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, middle.FailWithMsg("Error change bytes"))
		return
	}
	println(file.Filename)
	filePath := TSconf.temSaveBaseUrl + file.Filename
	if err := os.WriteFile(filePath, fileBytes, 0666); err != nil {
		c.JSON(http.StatusInternalServerError, middle.FailWithMsg("Error write file"))
	}
	c.JSON(http.StatusOK, middle.SuccessWithNil())
}
