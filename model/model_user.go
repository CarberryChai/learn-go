package model

import (
	"gorm.io/gorm"
)

//User 用户表模型
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}
