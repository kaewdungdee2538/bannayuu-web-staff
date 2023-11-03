package api

import (
	constants "bannayuu-web-admin/constants"
	controller_slot_add "bannayuu-web-admin/controllers/slot/add"
	controller_slot_get "bannayuu-web-admin/controllers/slot/get"
	interceptor_company "bannayuu-web-admin/interceptor/company"
	interceptor "bannayuu-web-admin/interceptor/jwt"
	interceptor_slot "bannayuu-web-admin/interceptor/slot"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupSlotAPI(router *gin.Engine) {
	slotApiHTTP := constants.GetSlotHTTPClient()
	fmt.Printf("slot api http : %s", slotApiHTTP)
	slotApi := router.Group(slotApiHTTP)
	{
		slotApi.POST("/add/manual",
			interceptor.JwtVerify,
			interceptor_slot.CheckAddSlotValueInterceptor,
			controller_slot_add.AddSlotManual)
		slotApi.POST("/get/not-use",
			interceptor.JwtVerify,
			interceptor_company.GetIdCompanyJsonValidateValuesInterceptor,
			controller_slot_get.GetSlotNotUse)
			slotApi.POST("/get/max",
				interceptor.JwtVerify,
				interceptor_company.GetIdCompanyJsonValidateValuesInterceptor,
				controller_slot_get.GetSlotMax)
	}
}
