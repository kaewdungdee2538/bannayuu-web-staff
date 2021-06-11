package interceptor

import (
	constants "bannayuu-web-admin/constants"
	home_intercep "bannayuu-web-admin/interceptor/home"
	user_model "bannayuu-web-admin/model/user"
	format_utls "bannayuu-web-admin/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func CheckGetUserValueInterceptor(c *gin.Context) {
	var userRequestModel user_model.UserAddRequestModel
	buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)

	err := json.Unmarshal([]byte(jsonString), &userRequestModel)

	if err != nil {
		//--------create error log
		format_utls.WriteLog(format_utls.GetErrorLogUserFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
		fmt.Printf("Error parsing JSON string - %s", err)
	}

	// if err := c.ShouldBind(&companyModel); err != nil {
	// 	c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	// 	c.Abort()
	// 	return
	// }
	fmt.Print(userRequestModel)
	isErr, msg := checkUserAddRequest(&userRequestModel)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		// ---------Convert obj to json string
		userInfo, err := json.Marshal(userRequestModel)
		if err != nil {
			format_utls.WriteLogInterface(format_utls.GetErrorLogUserFile(), nil, constants.MessageCovertObjTOJSONFailed)
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
			return
		}
		// -----forward request body middleware to endpoint
		rdr2 := ioutil.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf("%v", string(userInfo)))))
		c.Request.Body = rdr2
		c.Next()
	}
}
func checkUserAddRequest(userRequestModel *user_model.UserAddRequestModel) (bool, string) {
	first_name := userRequestModel.First_name
	last_name := userRequestModel.Last_name
	address := userRequestModel.Address
	mobile := strings.TrimSpace(userRequestModel.Mobile)
	line := userRequestModel.Line
	email := userRequestModel.Email
	username := strings.TrimSpace(userRequestModel.Username)
	password := strings.TrimSpace(userRequestModel.Password)
	employee_privilege_id := strings.TrimSpace(userRequestModel.Employee_privilege_id)
	status := strings.TrimSpace(userRequestModel.Status)
	employee_type := strings.TrimSpace(userRequestModel.Employee_type)
	company_id := strings.TrimSpace(userRequestModel.Company_id)
	if len(first_name) == 0 {
		return true, constants.MessageFirstNameNotFound
	} else if format_utls.IsNotStringAlphabetRemark(first_name) {
		return true, constants.MessageFirstNameProhitbitSpecial
	} else if len(last_name) == 0 {
		return true, constants.MessageLastNameNotFound
	} else if format_utls.IsNotStringAlphabetRemark(last_name) {
		return true, constants.MessageLastNameProhitbitSpecial
	} else if format_utls.IsNotStringAlphabetRemark(address) {
		return true, constants.MessageAddressProhibitSpecial
	} else if len(mobile) > 0 && len(mobile) != 10 {
		return true, constants.MessageMobileNotEqual10Character
	} else if format_utls.IsNotStringNumber(mobile) {
		return true, constants.MessageMobileNotNumber
	} else if format_utls.IsNotStringAlphabetRemark(line) {
		return true, constants.MessageLineProhibitSpecial
	} else if format_utls.IsNotStringAlphabetRemark(email) {
		return true, constants.MessageEmailFormatInValid
	} else if len(username) == 0 {
		return true, constants.MessageUsernameNotFount
	} else if format_utls.IsNotStringEngOrNumber(username) {
		return true, constants.MessageUsernameIsSpecialProhibit
	} else if len(password) == 0 {
		return true, constants.MessagePasswordNotFount
	} else if format_utls.IsNotStringEngOrNumber(password) {
		return true, constants.MessagePasswordIsSpecialProhibit
	} else if len(employee_privilege_id) == 0 {
		return true, constants.MessageEmployeePrivilegeIdNotFound
	} else if format_utls.IsNotStringNumber(employee_privilege_id) {
		return true, constants.MessageEmployeePrivilegeIdNotNumber
	} else if len(employee_type) == 0 {
		return true, constants.MessageEmployeeTypeNotFound
	}else if format_utls.IsNotStringAlphabet(employee_type){
		return true, constants.MessageEmployeeTypeProhibitSpecial
	}else if len(status) == 0 {
		return true, constants.MessageEmployeeStatusNotFound
	} else if format_utls.IsNotStringAlphabet(status) {
		return true, constants.MessageEmployeeStatusProhibitSpecial
	}
	errComId, msgComId := home_intercep.CheckValueCompanyIdNotDisavle(company_id)
	if errComId {
		return true, msgComId
	}
	return checkUserIsDuplicateInbaseWhenCreateUser(employee_type,username,company_id)
}
