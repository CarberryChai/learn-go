package controller

import (
	"net/http"

	"github.com/CarberryChai/learn-go/response"
	"github.com/gin-gonic/gin"
)

type registBind struct {
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required,gte=8"`
	ConfirmPassword string `form:"confirmPassword" binding:"requred,gte=8,eqfield=Password"`
}

func Register(ctx *gin.Context) {
	var userInput registBind
	if err := ctx.ShouldBind(&userInput); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(nil, err.Error()))
		return
	}
}
