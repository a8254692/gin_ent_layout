package router

import (
	"back-end/merchant/internal/logic"
	"back-end/merchant/pkg/bootstrap"
	"back-end/merchant/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, noCheckProRouter)
}

func noCheckProRouter(r *gin.RouterGroup, app *bootstrap.App) {

	user := logic.NewUser(app)
	v1 := r.Group("/")
	{
		v1.POST("/login", ginx.Handle(user.Login))
	}
}
