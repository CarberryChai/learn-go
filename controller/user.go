package controller

import (
	"net/http"

	"github.com/CarberryChai/learn-go/response"
	"github.com/CarberryChai/learn-go/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Register 注册
func Register(ctx *gin.Context) {
	var userInput service.RegistBind
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(nil, err.Error()))
		return
	}
	if res := (&userInput).Register(); res != nil {
		ctx.JSON(http.StatusOK, res)
		return
	}
	ctx.JSON(http.StatusOK, response.Success(nil, "创建成功"))
}

// Login 登录
func Login(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{"count": count})
}
