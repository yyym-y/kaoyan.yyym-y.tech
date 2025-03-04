package model

type PlayInfo struct {
	VedioId int    `gorm:"column:vedio_id;type:int;NOT NULL" json:"vedioId"`
	Num     int    `gorm:"column:num;type:int;NOT NULL" json:"uum"`
	Type    int    `gorm:"column:type;type:int;NOT NULL" json:"type"`
	Name    string `gorm:"column:name;type:varchar(255);" json:"name"`
	Url     string `gorm:"column:url;type:varchar(255);NOT NULL" json:"url"`
}

func (m *PlayInfo) TableName() string {
	return "playinfo"
}
