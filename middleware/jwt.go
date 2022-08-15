package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miaojiaxi2004/go_server/global"
	"github.com/miaojiaxi2004/go_server/model/common/response"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		fmt.Println(token)
		userData, err := global.CONFIG.JWT.JWTDecode(token)
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
			return
		}
		// 拿解析的数据做任何事情
		fmt.Println(userData)
		c.Next()
	}
}
