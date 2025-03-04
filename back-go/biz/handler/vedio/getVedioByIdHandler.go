package vedio

import (
	"back/biz/middle"
	"back/biz/model"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetVedioByIdHandler(ctx context.Context, c *app.RequestContext) {
	vedioId := c.Query("vedioId")
	res := model.Vedio{}
	if err := model.DB.Where("vedio_id = ?", vedioId).First(&res); err.Error != nil {
		c.JSON(http.StatusInternalServerError, middle.FailWithMsg("Error read vedioId"))
		return
	}
	c.JSON(http.StatusOK, middle.SuccessWithDate(res))
}
