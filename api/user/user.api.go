package api

import (
	constants "bannayuu-web-admin/constants"
	controller_user "bannayuu-web-admin/controllers/user"
	interceptor "bannayuu-web-admin/interceptor/jwt"
	interceptor_user "bannayuu-web-admin/interceptor/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupUserAPI(router *gin.Engine) {
	userApiHTTP := constants.GetUserHTTPClient()
	fmt.Printf("user api http : %s", userApiHTTP)
	authenApi := router.Group(userApiHTTP)
	{
		authenApi.POST("/create-user",
			interceptor.JwtVerify,
			interceptor_user.CheckGetUserValueInterceptor,
			controller_user.AddUser)
		authenApi.POST("/get-user",
			interceptor.JwtVerify,
			interceptor_user.CheckGetUserInfoValueInterceptor,
			controller_user.GetHomeAll)
		authenApi.POST("/get-userinfo-byid",
			interceptor.JwtVerify,
			interceptor_user.CheckGetUserInfoByIdValueInterceptor,
			controller_user.GetHomeInfo)
	}
}
