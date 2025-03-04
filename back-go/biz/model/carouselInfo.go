package model

type CarouselInfo struct {
	CarouselId  int    `gorm:"column:carousel_id;type:int;NOT NULL" json:"carouselId"`
	Cover       string `gorm:"column:cover;type:varchar(255);NOT NULL" json:"cover"`
	IfShowVedio bool   `gorm:"column:if_show_vedio;type:boolean;NOT NULL" json:"ifShowVedio"`
	VedioId     int    `gorm:"column:vedio_id;type:int" json:"vedioId"`
}

func (m *CarouselInfo) TableName() string {
	return "carousel_info"
}
