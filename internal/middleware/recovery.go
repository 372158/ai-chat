package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recovery 关卡：handler 里 panic 时的安全网
// defer + recover() 接住 panic，回 500 而不是整个进程崩溃（第 05 课的知识上岗）
// 注意：500 是 HTTP 状态码直接 500——传输层事故，和业务 fail() 的 200+code 不同
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "服务器内部错误",
				})
			}
		}()
		c.Next()
	}
}
