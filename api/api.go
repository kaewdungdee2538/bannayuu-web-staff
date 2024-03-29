package api

import (
	authen "bannayuu-web-admin/api/authen"
	company "bannayuu-web-admin/api/company"
	home "bannayuu-web-admin/api/home"
	user "bannayuu-web-admin/api/user"
	villager "bannayuu-web-admin/api/villager"
	slot_api "bannayuu-web-admin/api/slot"
	constants "bannayuu-web-admin/constants"
	"bannayuu-web-admin/db"
	interceptor "bannayuu-web-admin/interceptor/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	router.MaxMultipartMemory = 8 << 20 // 8mb
	db.SetupDB(constants.DbHost, constants.DbUserName, constants.DbPassword, constants.DbName, constants.DbPort, constants.DbMaxIdleTime, constants.DbMaxConnections)
	authen.SetupAuthenAPI(router)
	company.SetupCompanyAPI(router)
	home.SetupHomeAPI(router)
	villager.SetupVillagerAPI(router)
	user.SetupUserAPI(router)
	slot_api.SetupSlotAPI(router)
	setupTest(router)
}

func setupTest(router *gin.Engine) {
	authApiHTTP := constants.GetHTTPClient()
	router.GET(fmt.Sprintf("/%s", authApiHTTP), interceptor.JwtVerify, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"error": false,
			"result":  "Hello Staff",
			"message": "Hello Staff"})
	})
}
