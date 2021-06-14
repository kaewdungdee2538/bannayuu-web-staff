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
			controller_user.GetUserAll)
		authenApi.POST("/get-userinfo-byid",
			interceptor.JwtVerify,
			interceptor_user.CheckGetUserInfoByIdValueInterceptor,
			controller_user.GetHomeInfo)
		authenApi.POST("/edit-userinfo",
			interceptor.JwtVerify,
			interceptor_user.CheckGetUserValueWhenEditInfoInterceptor,
			controller_user.EditUser)
		authenApi.POST("/change-privilege",
			interceptor.JwtVerify,
			interceptor_user.CheckChangePrivilegeUserValidateValuesInterceptor,
			controller_user.ChangePrivilegeUser)
		authenApi.POST("/get-user-is-below-myself",
			interceptor.JwtVerify,
			interceptor_user.CheckGetUserInfoValueInterceptor,
			controller_user.GetUserIsBelowMyselfAll)
		authenApi.POST("/change-main-company",
			interceptor.JwtVerify,
			interceptor_user.CheckChangeMainCompanyUserValidateValuesInterceptor,
			controller_user.ChangeMainCompanyUser)
		authenApi.POST("/addordelete-company-list",
			interceptor.JwtVerify,
			interceptor_user.CheckAddOrDeleteCompanyListUserValidateValuesInterceptor,
			controller_user.AddOrDeleteCompanyListUser)
		authenApi.POST("/get-privilege",
			interceptor.JwtVerify,
			controller_user.GetPrivilege)
	}
}
