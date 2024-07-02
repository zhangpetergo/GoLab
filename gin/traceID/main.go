package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TraceIDMiddleware 是一个添加 traceID 到每个请求上下文的中间件
func TraceIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成一个新的 traceID
		traceID := uuid.New().String()

		// 将 traceID 添加到上下文
		c.Set("traceID", traceID)

		// 打印 traceID 以便调试
		c.Writer.Header().Set("traceID", traceID)

		// 继续处理请求
		c.Next()

		println(c.Writer.Status())
	}
}

func main() {
	r := gin.Default()

	// 使用 TraceIDMiddleware 中间件
	r.Use(TraceIDMiddleware())

	// 注册一些示例路由
	r.GET("/ping", func(c *gin.Context) {
		traceID, _ := c.Get("traceID")
		c.JSON(200, gin.H{
			"message": "pong",
			"traceID": traceID,
		})
	})

	r.POST("/submit", func(c *gin.Context) {
		traceID, _ := c.Get("traceID")
		c.JSON(200, gin.H{
			"status":  "submitted",
			"traceID": traceID,
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		traceID, _ := c.Get("traceID")
		c.JSON(200, gin.H{
			"user":    c.Param("id"),
			"traceID": traceID,
		})
	})

	r.Run("0.0.0.0:8888") // 启动服务

	errors.As()
}
