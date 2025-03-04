package index

import (
	"back/biz/middle"
	"back/biz/model"
	"context"
	"net/http"
	"sort"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetVedioTypeHandler(ctx context.Context, c *app.RequestContext) {
	result := make([]model.VedioType, 0)

	if err := model.DB.Find(&result); err.Error != nil {
		c.JSON(http.StatusInternalServerError, middle.FailWithMsg("Error read database"))
		return
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].TypeId < result[j].TypeId
	})
	c.JSON(http.StatusOK, middle.SuccessWithDate(result))
}
