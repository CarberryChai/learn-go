package router

import (
	"github.com/CarberryChai/learn-go/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.Static("/", "clients")
	r.POST("/upload", controller.Upload)
}
