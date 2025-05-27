package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 中间件用于记录API请求日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		t := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		
		// 处理请求
		c.Next()
		
		// 请求后
		latency := time.Since(t)
		statusCode := c.Writer.Status()
		
		log.Printf("[%s] | %3d | %12v | %s | %s",
			method, statusCode, latency, path, c.ClientIP())
	}
}

// CORS 中间件用于处理跨域请求
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
