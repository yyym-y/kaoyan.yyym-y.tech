package admin

import (
	"back/biz/middle"
	"back/biz/model"
	reqmodel "back/biz/reqModel"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func CreateCarouselHandler(ctx context.Context, c *app.RequestContext) {
	req := reqmodel.CreateCarouselReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, middle.FailWithMsg("bind wrong"))
		return
	}
	imgs, err := middle.ReadFile(req.Cover)
	if err != nil {
		c.JSON(400, middle.FailWithMsg("Cache find miss"))
		return
	}
	fileName := middle.NewUUID() + ".png"

	errChan := make(chan error)
	go func() {
		errChan <- writeImg(fileName, imgs)
	}()

	go func() {
		err := <-errChan
		if err != nil {
			c.JSON(500, middle.FailWithMsg("Failed to write image: "+err.Error()))
		} else {
			result := model.DB.Create(model.CarouselInfo{
				Cover:       fileName,
				IfShowVedio: req.IfShowVedio,
				VedioId:     req.VedioId,
			})
			if result.Error != nil {
				c.JSON(500, middle.FailWithMsg("Failed to write datwbase"))
			}
			c.JSON(http.StatusOK, middle.SuccessWithNil())
		}
	}()
}
