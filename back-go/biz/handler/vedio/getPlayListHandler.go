package vedio

import (
	"back/biz/middle"
	"back/biz/model"
	reqmodel "back/biz/reqModel"
	"context"
	"net/http"
	"sort"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetPlayListHandler(ctx context.Context, c *app.RequestContext) {
	req := reqmodel.VedioPlayInfoReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, middle.FailWithMsg("Error bind"))
		return
	}
	playInfos := make([]model.PlayInfo, 0)
	if err := model.DB.Where("vedio_id = ?", req.VedioId).Find(&playInfos); err.Error != nil {
		c.JSON(500, middle.FailWithMsg("Error read datebase"))
		return
	}
	sort.Slice(playInfos, func(i, j int) bool {
		return playInfos[i].Num < playInfos[j].Num
	})
	c.JSON(http.StatusOK, middle.SuccessWithDate(playInfos))
}
