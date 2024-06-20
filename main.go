package main

import (
	"back-end/driver"
	"back-end/merchant/cmd/api"
)

func main() {
	coreSrv := api.NewCoreService()
	driver.Drive(coreSrv)
}
