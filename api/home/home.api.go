package api

import (
	constants "bannayuu-web-admin/constants"
	controller_home "bannayuu-web-admin/controllers/home"
	interceptor "bannayuu-web-admin/interceptor/jwt"
	interceptor_home "bannayuu-web-admin/interceptor/home"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupHomeAPI(router *gin.Engine) {
	homeApiHTTP := constants.GetHomeHTTPClient()
	fmt.Printf("home api http : %s", homeApiHTTP)
	authenApi := router.Group(homeApiHTTP)
	{
		authenApi.POST("/import-array",
		 interceptor.JwtVerify,
		 interceptor_home.CheckAddHomeArrayValuesInterceptor,
		  controller_home.AddHomeArray)
	}
}
