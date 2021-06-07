package api

import (
	constants "bannayuu-web-admin/constants"
	controller_villager "bannayuu-web-admin/controllers/villager"
	interceptor "bannayuu-web-admin/interceptor/jwt"
	interceptor_villager "bannayuu-web-admin/interceptor/villager"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupVillagerAPI(router *gin.Engine) {
	villagerApiHTTP := constants.GetVillagerHTTPClient()
	fmt.Printf("villager api http : %s", villagerApiHTTP)
	authenApi := router.Group(villagerApiHTTP)
	{
		authenApi.POST("/import-array",
			interceptor.JwtVerify,
			interceptor_villager.CheckAddVillagerArrayValuesInterceptor,
			controller_villager.AddVillagerArray)
	}
}
