package logic

import (
	"back-end/merchant/internal/data/ent"
	"back-end/merchant/internal/data/ent/sysrolemenu"
	"back-end/merchant/internal/data/ent/sysrolepermissions"
	"back-end/merchant/internal/data/ent/sysroleusers"
	"back-end/merchant/pkg/bootstrap"
	"back-end/merchant/pkg/ginx"
	"back-end/merchant/pkg/statusx"
	"github.com/sirupsen/logrus"
)

func NewRole(app *bootstrap.App) *Role {
	return &Role{
		app,
	}
}

type Role struct {
	*bootstrap.App
}

func (s Role) GetList(c *ginx.Context) {
	list, err := s.Data.DB.SysRoles.Query().All(c)
	if err != nil {
		logrus.Errorf("Role GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s Role) GetMenuList(c *ginx.Context) {
	req := &ent.SysRoles{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role GetMenuList ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	list, err := s.Data.DB.SysRoleMenu.Query().Where(sysrolemenu.RoleID(req.ID)).All(c)
	if err != nil {
		logrus.Errorf("Role GetMenuList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s Role) GetUserList(c *ginx.Context) {
	req := &ent.SysRoles{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role GetUserList ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	list, err := s.Data.DB.SysRoleUsers.Query().Where(sysroleusers.RoleID(req.ID)).All(c)
	if err != nil {
		logrus.Errorf("Role GetUserList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s Role) GetPermissionsList(c *ginx.Context) {
	req := &ent.SysRoles{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role GetPermissionsList ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	list, err := s.Data.DB.SysRolePermissions.Query().Where(sysrolepermissions.RoleID(req.ID)).All(c)
	if err != nil {
		logrus.Errorf("Role GetPermissionsList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}
