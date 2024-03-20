package models

import "berjcode/dependency/common"

type User struct {
	ID           uint   `gorm:"primary_key"`
	UserName     string `gorm:"not null;type:nvarchar(60)"`
	NameSurname  string `gorm:"not null;type:nvarchar(100)"`
	Email        string `gorm:"not null;type:nvarchar(60)"`
	Salt         string `gorm:"not null;type:longtext"`
	PasswordHash string `gorm:"not null;type:longtext"`
	common.EntityBase
}
