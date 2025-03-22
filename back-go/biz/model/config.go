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
	LetexUAT      string `gorm:"column:latex_uat;type:varchar(255);NOT NULL" json:"latex_uat"`
	LetexAppId    string `gorm:"column:latex_app_id;type:varchar(255);NOT NULL" json:"latex_app_id"`
	LetexSecret   string `gorm:"column:latex_secret;type:varchar(255);NOT NULL" json:"latex_secret"`
	LatexLightApi string `gorm:"column:latex_light_api;type:varchar(255);NOT NULL" json:"latex_light_api"`
	LatexNomalApi string `gorm:"column:latex_nomal_api;type:varchar(255);NOT NULL" json:"latex_nomal_api"`
}

func (m *Config) TableName() string {
	return "config"
}
