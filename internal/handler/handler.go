package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// echoReq：请求体长什么样。反引号 json 标签 = 翻译条，把 JSON 里的 "message" 对到 Go 的 Message 字段
type echoReq struct {
	Message string	`json:"message"`
}

// ok：成功时用的回信模板。data 是 any（任意类型，随接口传什么数据都行）
func ok(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": data})
}


// fail：失败时用的回信模板。注意 HTTP 码还是 200——传输成功，业务错在 code
func fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg})
}

// Health：handler = 每个接口一个函数，参数永远是 c（工具腰带）
func Health(c *gin.Context) {
	ok(c, gin.H{"status": "up"})
}


func Echo(c *gin.Context) {
	var req echoReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fail(c, 1, err.Error())
		return
	}
	ok(c, gin.H{"echo": req.Message})
}

// RegisterRoutes：所有路由的「落户」都在这里登记
func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", Health)   // 来访 /health 的 GET 请求 → 交给 Health 处理
	r.POST("/echo", Echo)     // 来访 /echo 的 POST 请求 → 交给 Echo 处理
}