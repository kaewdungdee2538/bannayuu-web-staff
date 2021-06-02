package api

import (
	constants "bannayuu-web-admin/constants"
	controller_com "bannayuu-web-admin/controllers/company"
	interceptor_com "bannayuu-web-admin/interceptor/company"
	interceptor "bannayuu-web-admin/interceptor/jwt"
	interceptor_remark "bannayuu-web-admin/interceptor/remark"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupCompanyAPI(router *gin.Engine) {
	authApiHTTP := constants.GetCompanyInsertHTTPClient()
	fmt.Printf("comapny api http : %s", authApiHTTP)
	authenApi := router.Group(authApiHTTP)
	{
		authenApi.POST("/add", interceptor.JwtVerify, interceptor_com.AddCompanyValidateValuesInterceptor, controller_com.AddCompany)
		authenApi.POST("/edit-info",
			interceptor.JwtVerify,
			interceptor_com.EditCompanyValidateValuesInterceptor,
			controller_com.EditInfoCompany)
		authenApi.POST("/disable",
			interceptor.JwtVerify,
			interceptor_com.GetIdCompanyValidateValuesInterceptor,
			interceptor_remark.CheckRemarkValidateValueFormDataInterceptor,
			controller_com.DisableCompany)
		authenApi.POST("/get-all", interceptor.JwtVerify, controller_com.GetCompanyAll)
		authenApi.POST("/get-by-id", interceptor.JwtVerify, controller_com.GetCompanyById)

	}
}
