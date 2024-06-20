package logic

import (
	"back-end/common/libs/jwtHelper"
	"back-end/merchant/internal/data/ent"
	"back-end/merchant/internal/data/ent/user"
	"back-end/merchant/pkg/bootstrap"
	"back-end/merchant/pkg/ginx"
	"back-end/merchant/pkg/statusx"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func NewUser(app *bootstrap.App) *User {
	return &User{
		app,
	}
}

type User struct {
	*bootstrap.App
}

func (s User) Login(c *ginx.Context) {
	req := &ent.User{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("User Login ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	ps, err := s.Data.DB.User.Query().Where(user.Username(req.Username)).Only(c)
	if err != nil {
		logrus.Errorf("User Login DB Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	if ps == nil {
		logrus.Errorf("User Login Err: ps == nil")
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	//if util.MD5(req.Password) != ps.Password {
	//	c.Render(statusx.StatusInvalidRequest, nil)
	//	return
	//}

	now := time.Now()
	claims := jwtHelper.MerchantClaims{
		Id:           uint64(ps.ID),
		MerchantHost: "",
		MerchantId:   uint64(ps.MerchantID),
		TimeSpam:     int(now.Unix()),
		ExpiresTime:  now.Add(time.Hour * 24).Unix(),
		Role:         strings.Split(ps.Role, ","),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Hour * 24).Unix(),
			IssuedAt:  now.Unix(),
		},
	}
	token, err := jwtHelper.GenerateToken(claims, jwtHelper.DefaultSecretKey)
	if err != nil || token == "" {
		logrus.Errorf("User Login GenerateToken Err: %s", err)
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	ps.Password = ""
	ps.Role = ""
	data := map[string]interface{}{"user": ps, "token": token}
	c.RenderSuccess(data)
	return
}

func (s User) GetList(c *ginx.Context) {
	list, err := s.Data.DB.User.Query().All(c)
	if err != nil {
		logrus.Errorf("User GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}
