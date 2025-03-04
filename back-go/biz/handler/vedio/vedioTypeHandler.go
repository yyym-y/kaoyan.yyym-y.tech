package vedio

import (
	"back/biz/middle"
	"back/biz/model"
	respmodel "back/biz/respModel"
	"context"
	"net/http"
	"sort"

	"github.com/cloudwego/hertz/pkg/app"
)

func VedioTypeHandler(ctx context.Context, c *app.RequestContext) {
	result := respmodel.VedioTypeResp{}

	if err := model.DB.Find(&result.TypeInfos); err.Error != nil {
		c.JSON(http.StatusInternalServerError, middle.FailWithMsg("Error read database"))
		return
	}

	sort.Slice(result.TypeInfos, func(i, j int) bool {
		return result.TypeInfos[i].TypeId < result.TypeInfos[j].TypeId
	})
	c.JSON(http.StatusOK, middle.SuccessWithDate(result))
}
