package interceptor

import (
	constants "bannayuu-web-admin/constants"
	home_intercep "bannayuu-web-admin/interceptor/home"
	user_model "bannayuu-web-admin/model/user"
	format_utls "bannayuu-web-admin/utils"
	"fmt"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func CheckAddOrDeleteCompanyListUserValidateValuesInterceptor(c *gin.Context) {
	var userModel user_model.UserAddOrDeleteCompanyListRequestModel
	if err := c.ShouldBind(&userModel); err != nil {
		fmt.Printf("Combine Error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		c.Abort()
		return
	}
	fmt.Print(userModel)
	isErr, msg := checkValuesUserAddOrDeleteCompanyList(userModel)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
	
		c.Next()
	}
}

func checkValuesUserAddOrDeleteCompanyList(userModel user_model.UserAddOrDeleteCompanyListRequestModel) (bool, string) {
	company_id := strings.TrimSpace(userModel.Company_id)
	employee_id := strings.TrimSpace(userModel.Employee_id)
	remark := strings.TrimSpace(userModel.Remark)
	if len(company_id) == 0 {
		return true, constants.MessageOldCompanyIdNotFound
	} else if format_utls.IsNotStringNumber(company_id) {
		return true, constants.MessageOldCompanyIdNotNumber
	} else if len(employee_id) == 0 {
		return true, constants.MessageEmployeeIdNotFound
	} else if format_utls.IsNotStringNumber(employee_id) {
		return true, constants.MessageEmployeeIdNotNumber
	} else if len(remark) == 0 {
		return true, constants.MessageRemarkNotFount
	} else if format_utls.IsNotStringEngOrNumber(remark) {
		return true, constants.MessageRemarkProhibitSpecial
	} else if len(remark) < 10 {
		return true, constants.MessageRemarkIsLower10Character
	}
	errComId, msgComId := home_intercep.CheckValueCompanyIdNotDisavle(company_id)
	if errComId {
		return true, msgComId
	}
	errUserId, msgUserId := checkEmployeeIdInBase(employee_id, company_id)
	if errUserId {
		return true, msgUserId
	}
	return false, ""
}
