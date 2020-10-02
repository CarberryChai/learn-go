package main

import (
	"github.com/CarberryChai/learn-go/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	router.Init(app)
	app.Run(":3000")
}
