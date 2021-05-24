package main

import (
	"bannayuu-web-admin/api"
	logger "bannayuu-web-admin/interceptor/logger"
	constants "bannayuu-web-admin/constants"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("images", "./uploads/images")

	//---setup logfile
	constants.SetupOSPath();
	//---logger
	logger.CreateOrAppendLogger(router)
	//---initial router api
	api.Setup(router)
	router.Run(":4501")
}
