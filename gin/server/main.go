package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"text/tabwriter"
)

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// 使用 tabwriter 包来格式化输出
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)

	// 打印表头
	fmt.Fprintln(writer, "Method\tPath")

	// 打印所有已注册的路由
	for _, route := range r.Routes() {
		fmt.Fprintf(writer, "%s\t%s\n", route.Method, route.Path)
	}

	// 刷新 writer 缓冲区，确保所有内容都被写出
	writer.Flush()

	r.Run("0.0.0.0:8000")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		fmt.Println("Start timer")
		// Process request

		c.Next()
		// Stop timer
		fmt.Println("Stop timer")
	}
}

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		// Stop timer
		fmt.Println("Stop timer")
	}
}
