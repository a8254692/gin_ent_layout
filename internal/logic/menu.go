package logic

import (
	"back-end/merchant/internal/data/ent"
	"back-end/merchant/pkg/bootstrap"
	"back-end/merchant/pkg/ginx"
	"back-end/merchant/pkg/statusx"
	"github.com/sirupsen/logrus"
)

func NewMenu(app *bootstrap.App) *Menu {
	return &Menu{
		app,
	}
}

type Menu struct {
	*bootstrap.App
}

func (s Menu) GetList(c *ginx.Context) {
	list, err := s.Data.DB.SysMenu.Query().All(c)
	if err != nil {
		logrus.Errorf("Menu GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s Menu) Add(c *ginx.Context) {
	req := &ent.SysMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Menu Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	list, err := s.Data.DB.SysMenu.Query().All(c)
	if err != nil {
		logrus.Errorf("Menu GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s Menu) Edit(c *ginx.Context) {
	req := &ent.SysMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Menu Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	list, err := s.Data.DB.SysMenu.Query().All(c)
	if err != nil {
		logrus.Errorf("Menu GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}
