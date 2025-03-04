package reqmodel

type CreateCarouselReq struct {
	Cover       string `json:"cover"`
	IfShowVedio bool   `json:"ifShowVedio"`
	VedioId     int    `json:"vedioId"`
}
