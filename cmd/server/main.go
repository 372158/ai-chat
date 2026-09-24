package main

import (
	"github.com/372158/ai-chat/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()  		    // 造引擎（你造过的那个 Engine 的 Gin 版）

	handler.RegisterRouters(r)  // 把路由表挂上引擎
	r.Run(":8080")				// 开门营业，监听 8080
}