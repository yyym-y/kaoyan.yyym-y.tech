package model

type VedioType struct {
	TypeId int    `gorm:"column:type_id;type:INT;" json:"typeId"`
	Name   string `gorm:"column:name;type:VARCHAR(255);NOT NULL" json:"name"`
}

func (m *VedioType) TableName() string {
	return "vedio_type"
}
