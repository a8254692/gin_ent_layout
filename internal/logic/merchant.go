package logic

import (
	"back-end/merchant/pkg/bootstrap"
	"back-end/merchant/pkg/ginx"
	"back-end/merchant/pkg/statusx"
	"github.com/sirupsen/logrus"
)

func NewMerchant(app *bootstrap.App) *Merchant {
	return &Merchant{
		app,
	}
}

type Merchant struct {
	*bootstrap.App
}

func (s Merchant) GetList(c *ginx.Context) {
	list, err := s.Data.DB.MerchantList.Query().All(c)
	if err != nil {
		logrus.Errorf("Merchant GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}
