package main

import (
	"bannayuu-web-admin/api"
	constants "bannayuu-web-admin/constants"
	logger "bannayuu-web-admin/interceptor/logger"
	utils "bannayuu-web-admin/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	//---get data from environment
	if constants.InitializeEnv() {
		//---setup static pah
		router.Static("images", constants.RootImages)
		//---setup CORS
		router.Use(CORSMiddleware())
		//---setup logfile
		constants.SetupOSPath()
		//---logger
		logger.CreateOrAppendLogger(router)
		//---initial router api
		api.Setup(router)
		router.Run(":"+constants.AppPort)
	}else{
		utils.WriteLog(utils.GetErrorLogFile(), fmt.Sprintln("Get Environment Fail !!"))
		fmt.Println("Get Environment Fail !!")
		return
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
