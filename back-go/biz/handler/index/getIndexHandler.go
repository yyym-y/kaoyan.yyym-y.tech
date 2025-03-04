package index

import (
	"back/biz/middle"
	"back/biz/model"
	"context"
	"net/http"
	"sort"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetIndexHandler(ctx context.Context, c *app.RequestContext) {
	typeInfos := make([]model.VedioType, 0)
	if err := model.DB.Find(&typeInfos); err.Error != nil {
		c.JSON(http.StatusInternalServerError, middle.FailWithMsg("Error read database"))
		return
	}
	sort.Slice(typeInfos, func(i, j int) bool {
		return typeInfos[i].TypeId < typeInfos[j].TypeId
	})
	mp := make(map[int][]model.Vedio)

	for i := 0; i < len(typeInfos); i++ {
		var videos []model.Vedio
		if err := model.DB.Where("type_id = ?", i+1).Where("visable = true").Find(&videos).Limit(100); err.Error != nil {
			c.JSON(http.StatusInternalServerError, middle.FailWithMsg("Error read database"))
		}
		mp[typeInfos[i].TypeId] = videos
	}
	c.JSON(http.StatusOK, middle.SuccessWithDate(mp))
}
