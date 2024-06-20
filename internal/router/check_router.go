package router

import (
	"back-end/merchant/internal/logic"
	"back-end/merchant/pkg/bootstrap"
	"back-end/merchant/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, checkProRouter)
}

func checkProRouter(r *gin.RouterGroup, app *bootstrap.App, gHanFun ...gin.HandlerFunc) {
	v1 := r.Group("/").Use(gHanFun...)
	{
		user := logic.NewUser(app)
		menu := logic.NewMenu(app)
		merchant := logic.NewMerchant(app)
		role := logic.NewRole(app)
		permission := logic.NewPermissions(app)

		v1.GET("/user", ginx.Handle(user.GetList))
		v1.GET("/menu", ginx.Handle(menu.GetList))
		v1.GET("/role", ginx.Handle(role.GetList))
		v1.GET("/role/user", ginx.Handle(role.GetUserList))
		v1.GET("/role/menu", ginx.Handle(role.GetMenuList))
		v1.GET("/role/permission", ginx.Handle(role.GetPermissionsList))
		v1.GET("/merchant", ginx.Handle(merchant.GetList))
		v1.GET("/permission", ginx.Handle(permission.GetList))
	}
}
