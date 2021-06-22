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
	companyApiHTTP := constants.GetCompanyInsertHTTPClient()
	fmt.Printf("comapny api http : %s", companyApiHTTP)
	authenApi := router.Group(companyApiHTTP)
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
		authenApi.POST("/enable",
			interceptor.JwtVerify,
			interceptor_com.GetIdCompanyValidateValuesInterceptor,
			interceptor_remark.CheckRemarkValidateValueFormDataInterceptor,
			controller_com.EnableCompany)
		authenApi.POST("/get-all", interceptor.JwtVerify, controller_com.GetCompanyAll)
		authenApi.POST("/get-all-not-disable", interceptor.JwtVerify, controller_com.GetCompanyAllNotDisable)
		authenApi.POST("/get-by-id", interceptor.JwtVerify, controller_com.GetCompanyById)
		authenApi.POST("/get-companylist-all", interceptor.JwtVerify, controller_com.GetCompanyListAll)
		authenApi.POST("/get-companylist-all-not-cit-company", interceptor.JwtVerify, controller_com.GetCompanyAllIsNotCitCompany)
		
	}
}
