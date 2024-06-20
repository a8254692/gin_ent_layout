package router

import (
	"back-end/merchant/internal/router/middleware"
	"back-end/merchant/pkg/bootstrap"
	"github.com/gin-gonic/gin"
)

// InitRouter 路由初始化
func InitRouter(app *bootstrap.App) *gin.Engine {
	r := gin.New()
	if r == nil {
		panic("load gin error")
	}

	gin.SetMode(gin.DebugMode)

	// 自定义错误处理
	r.Use(middleware.CustomError)
	// NoCache is a middleware function that appends headers
	r.Use(middleware.NoCache)
	// 跨域处理
	r.Use(middleware.Options)

	// the jwt middleware
	authMid := middleware.JWTAuth()

	// 注册业务路由
	ExecRouter(r, app, authMid)

	return r
}
