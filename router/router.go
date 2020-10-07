package router

import (
	"github.com/CarberryChai/learn-go/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Init 路由注册
func Init(r *gin.Engine) {
	//	r.Static("/", "public")
	store := cookie.NewStore([]byte("lovelove123"))
	router := r.Group("/api")
	router.POST("/upload", controller.Upload)
	router.POST("/register", controller.Register)
	router.POST("/login", sessions.Sessions("uid", store), controller.Login)
}
