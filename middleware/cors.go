package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORS 中间件，处理跨域请求
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许的来源，这里设置为*表示允许所有来源
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 设置允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// 设置是否允许发送Cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 设置预检请求的有效期
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// 确保在所有响应中都包含这些头信息
		c.Writer.Header().Set("Vary", "Origin")

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // 204 No Content
			return
		}

		c.Next()
	}
}
