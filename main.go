package main

import (
	"test/config"
	"test/controller"
)

// @title           --Swagger Example API
// @version         1234.0
// @license.name Apache 2.0
// @host      http://api-love-dev-zxm.qingteng-inc.com:6190
// @securityDefinitions.basic  BasicAuth
func main() {
	config.InitConfig("./config")
	router := controller.Router()
	router.Run(":6190")
}
