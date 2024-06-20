/*
 * @Date: 2024-05-20 11:40:47
 * @LastEditTime: 2024-05-22 10:06:31
 * @Description:
 */
package api

import (
	"back-end/driver/structs"
	"back-end/driver/structs/service"

	"github.com/sirupsen/logrus"
)

type merchantCore struct {
	e *structs.Exposer
}

// NewCoreService 创建服务
func NewCoreService() service.Service {
	return new(merchantCore)
}

// InitServer 初始化服务
func InitServer() error {

	Run()

	return nil
}

func (c *merchantCore) Init(e *structs.Exposer, param ...string) error {
	err := InitServer()
	if err != nil {
		logrus.Errorf("server init failed,err:(%v)", err)
		return err
	}

	logrus.Infof("server init succeed ...")

	return nil
}

func (c *merchantCore) Start() error {
	return nil
}
