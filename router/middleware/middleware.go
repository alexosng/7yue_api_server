package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/handler"
	"github.com/Away0x/7yue_api_server/constant/errno"
	)

// 跨域
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

// appkey
// header 携带的优先级更高
func KeyAuth(c *gin.Context) {
	if c.Request.URL.String() == "/swagger/index.html" {
		c.Next()
		return
	}

	key := c.Request.Header.Get("appkey")
	if key == "" {
		key = c.Query("appkey")
	}

	// 没有 key
	if key != "abc" {
		handler.SendResponse(c, errno.New(errno.AuthError, nil).Addf("appkey 错误 - %s", key), nil)
		c.Abort()
		return
	}
	c.Next()
}