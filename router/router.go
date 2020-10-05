package router

import (
	"github.com/CarberryChai/learn-go/controller"
	"github.com/gin-gonic/gin"
)

// Init 路由注册
func Init(r *gin.Engine) {
	//	r.Static("/", "public")
	r.POST("/upload", controller.Upload)
}
