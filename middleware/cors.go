package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORS 中间件，处理跨域请求(浏览器安全机制)
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义允许访问的域名白名单
		allowedOrigins := []string{
			"http://127.0.0.1:8000",
			"http://localhost:8000", // 也可以同时添加localhost:8000
		}
		// 获取请求来源
		origin := c.Request.Header.Get("Origin")
		// 检查请求来源是否在白名单中
		allowed := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				allowed = true
				break
			}
		}
		// 如果在白名单中，则设置允许跨域
		if allowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// 设置允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") //  PUT, DELETE,
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

		// 传递下一个请求
		c.Next()
	}
}
