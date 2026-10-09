package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// jwtSecret 签名密钥（HS256 是对称加密：签发和验票用同一把钥匙）
// TODO: 冲刺结束后挪进 config 包，从环境变量读，不写死在代码里
const jwtSecret = "dev-secret"

// GenToken 登录成功后的发票：payload 装 user_id，带过期时间（Unix 时间戳）
func GenToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// Auth 验票关卡：从请求头 Authorization: Bearer <token> 取票验票
// 没带票 / 假票 / 过期票 → 401 拦在门外（Abort：后面的路不许走）
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "未授权：缺少 token"})
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			return []byte(jwtSecret), nil // 验票用同一把钥匙
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "未授权：token 无效"})
			return
		}

		// 验过了：把 user_id 存进上下文，handler 直接取用，不用再解析一遍
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
		}
		c.Next()
	}
}
