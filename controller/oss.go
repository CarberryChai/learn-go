package controller

import (
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/CarberryChai/learn-go/response"
	"github.com/gin-gonic/gin"
)

type BindFile struct {
	Name string                `form:"name" binding:"required"`
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func Upload(ctx *gin.Context) {
	var uploadFile BindFile
	if err := ctx.ShouldBind(&uploadFile); err != nil {
		ctx.String(http.StatusBadRequest, "error: %s", err.Error())
		return
	}
	file := uploadFile.File
	dst := filepath.Base(file.Filename)
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.String(http.StatusBadRequest, "upload file error: %s", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response.Success(nil, "success"))
}
