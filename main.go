package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// app := gin.Default()
	//	router.Init(app)
	//	app.Run(":3000")
	name := os.Getenv("NAME")
	fmt.Println(name)
}
