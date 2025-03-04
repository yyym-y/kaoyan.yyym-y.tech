package model

type User struct {
	Username string `gorm:"column:username;type:varchar(255);NOT NULL" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`
	Email    string `gorm:"column:email;type:varchar(255);NOT NULL" json:"email"`
	Type     int    `gorm:"column:type;type:int;NOT NULL" json:"type"`
}

func (m *User) TableName() string {
	return "user"
}
