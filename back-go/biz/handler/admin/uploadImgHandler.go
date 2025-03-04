package admin

import (
	"back/biz/middle"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func UploadImgHandler(ctx context.Context, c *app.RequestContext) {
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
	err = middle.SaveFile(file.Filename, fileBytes)
	if err != nil {
		c.JSON(500, middle.FailWithMsg("cache fail"))
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, middle.SuccessWithNil())
}
