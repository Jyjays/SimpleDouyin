package models

type User struct {
	Id       int64  `gorm:"primary_key"`
	Username string `gorm:"size:32;primary_key"`
	Password string `gorm:"size:32;notnull"`
}
