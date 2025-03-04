package respmodel

import "back/biz/model"

type CarouselListResq struct {
	CarouselId  int         `json:"carouselId"`
	Cover       string      `json:"cover"`
	IfShowVedio bool        `json:"ifShowVedio"`
	Vedio       model.Vedio `json:"vedio"`
}
