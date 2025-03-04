package index

import (
	"back/biz/middle"
	"back/biz/model"
	respmodel "back/biz/respModel"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetCarouselHandler(ctx context.Context, c *app.RequestContext) {
	carousels := make([]model.CarouselInfo, 0)
	if err := model.DB.Find(&carousels); err.Error != nil {
		c.JSON(http.StatusInternalServerError, middle.FailWithMsg("Error read database"))
		return
	}
	res := make([]respmodel.CarouselListResq, 0)

	for i := 0; i < len(carousels); i++ {
		tem := respmodel.CarouselListResq{
			CarouselId:  carousels[i].CarouselId,
			Cover:       carousels[i].Cover,
			IfShowVedio: carousels[i].IfShowVedio,
		}
		if !carousels[i].IfShowVedio {
			res = append(res, tem)
			continue
		}
		var vedio model.Vedio
		if err := model.DB.Where("vedio_id = ?", carousels[i].VedioId).Find(&vedio); err.Error != nil {
			c.JSON(http.StatusInternalServerError, middle.FailWithMsg("Error read database"))
		}
		tem.Vedio = vedio
		res = append(res, tem)
	}
	c.JSON(http.StatusOK, middle.SuccessWithDate(res))
}
