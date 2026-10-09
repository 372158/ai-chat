package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 关卡：每个请求花多久、结果如何，记一笔
// 中间件 canonical 模式：进门前按秒表 → c.Next() 放行 → 出门后算总账
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // 进门前：秒表按下

		c.Next() // 放行！请求往里走（后面的关卡 + handler）

		// 出门后：Next() 返回才会执行到这——此时 handler 早干完了
		cost := time.Since(start)
		println(c.Request.Method, c.Request.URL.Path, c.Writer.Status(), cost.String())
	}
}
