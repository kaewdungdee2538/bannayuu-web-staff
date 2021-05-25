package api
import (
	constants "bannayuu-web-admin/constants"
	interceptor "bannayuu-web-admin/interceptor/jwt"
	interceptor_com "bannayuu-web-admin/interceptor/company"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupCompanyAPI(router *gin.Engine) {
	authApiHTTP := constants.GetCompanyInsertHTTPClient()
	fmt.Printf("comapny api http : %s", authApiHTTP)
	authenApi := router.Group(authApiHTTP)
	{
		authenApi.POST("/add", interceptor.JwtVerify,interceptor_com.AddCompanyValidateValues,AddCompany)
	}
}