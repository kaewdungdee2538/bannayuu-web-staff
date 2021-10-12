package interceptor

import (
	constants "bannayuu-web-admin/constants"
	home_intercep "bannayuu-web-admin/interceptor/home"
	user_model "bannayuu-web-admin/model/user"
	format_utls "bannayuu-web-admin/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckResetPasswordUserValueInterceptor(c *gin.Context) {
	var userRequestModel user_model.UserResetPasswordRequestModel
	buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)

	err := json.Unmarshal([]byte(jsonString), &userRequestModel)

	if err != nil {
		//--------create error log
		format_utls.WriteLog(format_utls.GetErrorLogUserResetPasswordFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
		fmt.Printf("Error parsing JSON string - %s", err)
	}

	// if err := c.ShouldBind(&companyModel); err != nil {
	// 	c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	// 	c.Abort()
	// 	return
	// }
	fmt.Print(userRequestModel)
	isErr, msg := checkResetPasswordUserRequest(&userRequestModel)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		// ---------Convert obj to json string
		userInfo, err := json.Marshal(userRequestModel)
		if err != nil {
			format_utls.WriteLogInterface(format_utls.GetErrorLogUserResetPasswordFile(), nil, constants.MessageCovertObjTOJSONFailed)
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
			return
		}
		// -----forward request body middleware to endpoint
		rdr2 := ioutil.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf("%v", string(userInfo)))))
		c.Request.Body = rdr2
		c.Next()
	}
	
}
func checkResetPasswordUserRequest(userRequestModel *user_model.UserResetPasswordRequestModel) (bool, string) {
	employee_id := strings.TrimSpace(userRequestModel.Employee_id)
	remark := strings.TrimSpace(userRequestModel.Remark)
	hold_time := strings.TrimSpace(userRequestModel.Hold_time)
	company_id := strings.TrimSpace(userRequestModel.Company_id)
	if len(employee_id) == 0 {
		return true, constants.MessageEmployeeIdNotFound
	} else if format_utls.IsNotStringNumber(employee_id) {
		return true, constants.MessageEmployeeIdNotNumber
	} else if len(remark) == 0 {
		return true, constants.MessageRemarkNotFount
	} else if format_utls.IsNotStringAlphabetRemark(remark) {
		return true, constants.MessageRemarkProhibitSpecial
	} else if len(hold_time) == 0 {
		return true, constants.MessageHoldTimeNotFound
	} else if format_utls.IsNotStringAlphabet(hold_time) {
		return true, constants.MessageHoldTimeIsProhibitSpecial
	} else if errComId, msgComId := home_intercep.CheckValueCompanyIdNotDisavle(company_id); errComId {
		return true, msgComId
	} else if errEmp, msgEmp := checkEmployeeIdInBase(employee_id, company_id); errEmp {
		return true, msgEmp
	}
	return false, ""
}
