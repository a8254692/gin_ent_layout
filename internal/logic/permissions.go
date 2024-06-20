package logic

import (
	"back-end/merchant/pkg/bootstrap"
	"back-end/merchant/pkg/ginx"
	"back-end/merchant/pkg/statusx"
	"github.com/sirupsen/logrus"
)

func NewPermissions(app *bootstrap.App) *Permissions {
	return &Permissions{
		app,
	}
}

type Permissions struct {
	*bootstrap.App
}

func (s Permissions) GetList(c *ginx.Context) {
	list, err := s.Data.DB.SysPermissions.Query().All(c)
	if err != nil {
		logrus.Errorf("Permissions GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}
