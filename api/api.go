package api

import (
	authen "bannayuu-web-admin/api/authen"
	company "bannayuu-web-admin/api/company"
	constants "bannayuu-web-admin/constants"
	"bannayuu-web-admin/db"
	interceptor "bannayuu-web-admin/interceptor/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	db.SetupDB()
	authen.SetupAuthenAPI(router);
	company.SetupCompanyAPI(router);
	setupTest(router);
}

func setupTest(router *gin.Engine){
	authApiHTTP := constants.GetHTTPClient();
	router.GET(fmt.Sprintf("/%s", authApiHTTP), interceptor.JwtVerify, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"error": false,
			"result":  "Hello Staff",
			"message": "Hello Staff"})
	})
}
