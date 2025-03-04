package model

type Vedio struct {
	VedioId     int    `gorm:"column:vedio_id;type:int;NOT NULL" json:"vedioId"`
	VedioName   string `gorm:"column:vedio_name;type:varchar(255);NOT NULL" json:"vedioName"`
	VedioCover  string `gorm:"column:vedio_cover;type:varchar(255);NOT NULL" json:"vedioCover"`
	Description string `gorm:"column:description;type:varchar(255);" json:"description"`
	Visable     bool   `gorm:"column:visable;type:boolean;NOT NULL" json:"visable"`
	TypeId      int    `gorm:"column:type_id;type:INT;" json:"typeId"`
}

func (m *Vedio) TableName() string {
	return "vedio"
}
