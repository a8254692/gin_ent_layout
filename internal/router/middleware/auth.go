package middleware

import (
	"back-end/common/pkg"
	"back-end/merchant/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"

	"back-end/common/libs/jwtHelper"
	"back-end/merchant/pkg/ginx"
	"back-end/merchant/pkg/statusx"
)

// JWTAuth jwt验证
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(config.AUTHORIZATIONKEY)
		if token == "" {
			ginx.NewContext(c).Render(statusx.StatusUnauthorized, "nil", http.StatusUnauthorized)
			c.Abort()
			return
		}

		claims := &jwtHelper.MerchantClaims{}
		err := jwtHelper.ParseToken(token, claims, jwtHelper.DefaultSecretKey)
		if err == jwtHelper.ErrInvalidToken {
			ginx.NewContext(c).Render(statusx.StatusVerifyTokenError, "nil", http.StatusUnauthorized)
			c.Abort()
			return
		}

		// 检查 token 是否已经过期
		if int(claims.ExpiresTime) < pkg.Generate.NowInt() {
			ginx.NewContext(c).Render(statusx.StatusVerifyTokenError, "nil", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set(config.AUTHUSERKEY, claims)
		c.Request = c.Request.WithContext(c)
		c.Next()
	}
}
