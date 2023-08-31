package routes

import (
	"chatgpt/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() {

	r := gin.Default()
	r.Use(corsMiddleware())
	r.POST("ask", api.Ask)
	r.POST("create", api.Create)
	r.POST("delete", api.Delete)
	r.POST("detail", api.Detail)
	r.POST("getlist", api.GetList)
	r.POST("update", api.Update)

	r.Run(":8877")
}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 设置允许所有域名的请求
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") // 允许的 HTTP 方法
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type") // 允许的请求头

		// 处理 OPTIONS 预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
