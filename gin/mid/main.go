package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
\
func main() {
	r := gin.Default()

	// 注册中间件
	r.Use(middlewareA)
	r.Use(middlewareB)

	// 设置路由
	r.GET("/example", exampleHandler)

	r.Run("0.0.0.0:6666")
}

func middlewareA(c *gin.Context) {
	fmt.Println("Executing middleware A")
	c.Next() // 继续到下一个中间件或处理程序
	fmt.Println("Finished middleware A")
}

func middlewareB(c *gin.Context) {
	fmt.Println("Executing middleware B")
	c.Next() // 继续到下一个中间件或处理程序
	fmt.Println("Finished middleware B")
}

func exampleHandler(c *gin.Context) {
	fmt.Println("Executing example handler")
	c.String(200, "Hello, World!")
}
