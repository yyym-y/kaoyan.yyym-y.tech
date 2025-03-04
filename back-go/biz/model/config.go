package model

type Config struct {
	EmailFrom     string `gorm:"column:emailFrom;type:varchar(255);NOT NULL" json:"emailFrom"`
	EmailPassword string `gorm:"column:emailPassword;type:varchar(255);NOT NULL" json:"emailPassword"`
	EmailHost     string `gorm:"column:emailHost;type:varchar(255);NOT NULL" json:"emailHost"`
	ImgPath       string `gorm:"column:img_path;type:varchar(255);NOT NULL" json:"img_path"`
	VedioPath     string `gorm:"column:vedio_path;type:varchar(255);NOT NULL" json:"vedio_path"`
	TemSavePath   string `gorm:"column:tem_save_path;type:varchar(255);NOT NULL" json:"tem_save_path"`
	NotePath      string `gorm:"column:note_path;type:varchar(255);NOT NULL" json:"note_path"`
	LogPath       string `gorm:"column:log_path;type:varchar(255);NOT NULL" json:"log_path"`
}

func (m *Config) TableName() string {
	return "config"
}
