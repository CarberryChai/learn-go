package service

import (
	"github.com/CarberryChai/learn-go/model"
	"github.com/CarberryChai/learn-go/response"
	"github.com/CarberryChai/learn-go/utils"
)

//RegistBind 用户上传数据绑定
type RegistBind struct {
	Name            string `json:"name" binding:"required`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,gte=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,gte=8,eqfield=Password"`
}

// Register service 注册
func (r *RegistBind) Register() *response.Result {
	password, _ := utils.HashPassword(r.Password)
	user := model.User{Name: r.Name, Email: r.Email, Password: password}
	var count int64
	model.DB.Model(&user).Where("name = ?", r.Name).Count(&count)
	if count > 0 {
		return response.New(-401, nil, "该用户已存在，请重新注册")
	}
	result := model.DB.Create(&user)
	if result.Error != nil {
		return response.New(-201, nil, "创建用户失败，请稍后再试")
	}
	return nil
}
