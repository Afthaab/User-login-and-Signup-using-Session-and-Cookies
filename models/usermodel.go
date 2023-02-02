package models

import (
	"github.com/loginpage/config"
	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func NewUserModel() *UserModel {
	conn:= config.DBConn()

	return &UserModel{
		db: conn,
	}
}
