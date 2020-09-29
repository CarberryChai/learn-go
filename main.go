package main

import "fmt"

// import "github.com/gin-gonic/gin"

func main() {
	fmt.Println("hello, 世界！")
	z := add(1, 2)
	fmt.Println(z)
}

func add(x, y int) int {
	return x + y
}
