package api

import (
	"back-end/merchant/internal/router"
	"fmt"
	"github.com/spf13/viper"

	"back-end/merchant/internal/data"
	"back-end/merchant/pkg/bootstrap"
)

func Run() {
	// 定义一个命令行Flag来指定配置文件路径
	//path := pflag.StringP("config", "c", "", "Path to the config file")
	//pflag.Parse()
	//
	//err := bootstrap.LoadConfig(path)
	//if err != nil {
	//	panic("load config error: " + err.Error())
	//}

	//log, err := bootstrap.NewLogrusLogger()
	//if err != nil {
	//	panic("load log error: " + err.Error())
	//}

	dataSource, dFun, err := data.NewData()
	if err != nil {
		panic("load data error: " + err.Error())
	}
	defer dFun()

	app := &bootstrap.App{
		//Logger: log,
		Data: dataSource,
	}

	r := router.InitRouter(app)

	httpHost := viper.GetString("rpc_addr")
	httpPort := viper.GetInt("rpc_port")
	if httpHost == "" || httpPort <= 0 {
		panic("http path error: ")
	}

	err = r.Run(fmt.Sprintf("%s:%d", httpHost, httpPort))
	if err != nil {
		panic("http start error: " + err.Error())
	}

	return
}
