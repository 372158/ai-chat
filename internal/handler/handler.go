package handler

import (
	"net/http"

	"github.com/372158/ai-chat/internal/middleware"
	"github.com/gin-gonic/gin"
)

// echoReq：请求体长什么样。反引号 json 标签 = 翻译条，把 JSON 里的 "message" 对到 Go 的 Message 字段
type echoReq struct {
	Message string `json:"message"`
}

// loginReq：登录请求体
type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

// Login 换票处：校验身份 → 发票
// TODO: W2 接 MySQL 后，假校验换成查 users 表 + bcrypt 比密码
func Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, 1, err.Error())
		return
	}
	if req.Username != "酷米" || req.Password != "123456" {
		fail(c, 1001, "用户名或密码错误")
		return
	}
	token, err := middleware.GenToken(1) // userID 先写死 1，W2 换成查出来的真实 ID
	if err != nil {
		fail(c, 500, "签发 token 失败")
		return
	}
	ok(c, gin.H{"token": token})
}

// Profile 受保护测试接口：能走到这说明票已验过（Auth 关卡放行的）
func Profile(c *gin.Context) {
	userID, _ := c.Get("user_id") // Auth 关卡存进来的
	ok(c, gin.H{"user_id": userID})
}

// RegisterRoutes：所有路由的「落户」都在这里登记
func RegisterRoutes(r *gin.Engine) {
	// 全局关卡：所有请求都穿（Logger 记账、Recovery 保命）
	r.Use(middleware.Logger(), middleware.Recovery())

	// 公开路由：换票处不查票
	r.POST("/login", Login)

	// 路由组 /api：组内接口统一挂 Auth 关卡查票
	api := r.Group("/api")
	api.Use(middleware.Auth())
	api.GET("/profile", Profile)

	r.GET("/health", Health)
	r.POST("/echo", Echo)
}
