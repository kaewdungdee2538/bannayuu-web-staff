package interceptor

import (
	constants "bannayuu-web-admin/constants"
	user_model "bannayuu-web-admin/model/user"
	format_utls "bannayuu-web-admin/utils"
	home_intercep "bannayuu-web-admin/interceptor/home"
	// "bytes"
	// "encoding/json"
	// "io/ioutil"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckChangePrivilegeUserValidateValuesInterceptor(c *gin.Context) {
	var userModel user_model.UserChangePrivilegeRequestModel
	if err := c.ShouldBind(&userModel); err != nil {
		fmt.Printf("Combine Error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		c.Abort()
		return
	}
	fmt.Print(userModel)
	// buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	// jsonString := string(buf)

	// err := json.Unmarshal([]byte(jsonString), &companyModel)
	// if err != nil {
	// 	//--------create error log
	// 	format_utls.WriteLog(format_utls.GetErrorLogFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
	// 	fmt.Printf("Error parsing JSON string - %s", err)
	// }
	isErr, msg := checkValuesUserChangePrivilege(userModel)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		//---------Convert obj to json string
		// setup_data_map := map[string]interface{}{
		// 	"Company_code":               companyModel.Calculate_enable,
		// 	"except_time_split_from_day": companyModel.Except_time_split_from_day,
		// 	"price_of_cardloss":          companyModel.Price_of_cardloss}
		// err, _ := format_utls.ConvertInterfaceToJSON(setup_data_map)
		// if err {
		// 	c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
		// 	return
		// }
		//-----forward request body middleware to endpoint
		// rdr2 := ioutil.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf("%v", setup_data))))
		// c.Request.Body = rdr2
		c.Next()
	}
}

func checkValuesUserChangePrivilege(userModel user_model.UserChangePrivilegeRequestModel) (bool, string) {
	company_id := strings.TrimSpace(userModel.Company_id)
	employee_id := strings.TrimSpace(userModel.Employee_id)
	remark := strings.TrimSpace(userModel.Remark)
	employee_privilege_id := strings.TrimSpace(userModel.Employee_privilege_id)
	employee_type := strings.TrimSpace(userModel.Employee_type)
	if len(company_id) == 0 {
		return true, constants.MessageCompanyIdNotFound
	} else if format_utls.IsNotStringNumber(company_id) {
		return true, constants.MessageCompanyIdNotNumber
	} else if len(employee_id) == 0 {
		return true, constants.MessageEmployeeIdNotFound
	} else if format_utls.IsNotStringNumber(employee_id) {
		return true, constants.MessageEmployeeIdNotNumber
	} else if len(remark) == 0 {
		return true, constants.MessageRemarkNotFount
	} else if format_utls.IsNotStringAlphabetRemark(remark) {
		return true, constants.MessageRemarkProhibitSpecial
	} else if len(remark) < 10 {
		return true, constants.MessageRemarkIsLower10Character
	} else if len(employee_privilege_id) == 0 {
		return true, constants.MessageEmployeePrivilegeIdNotFound
	} else if format_utls.IsNotStringNumber(employee_privilege_id) {
		return true, constants.MessageEmployeePrivilegeIdNotNumber
	} else if len(employee_type) == 0 {
		return true, constants.MessageEmployeeTypeNotFound
	} else if format_utls.IsNotStringAlphabet(employee_type) {
		return true, constants.MessageEmployeeTypeProhibitSpecial
	}
	errComId, msgComId := home_intercep.CheckValueCompanyIdNotDisavle(company_id)
	if errComId {
		return true, msgComId
	}
	errUserId, msgUserId := checkEmployeeIdInBase(employee_id,company_id)
	if errUserId {
		return true, msgUserId
	}
	return checkEmployeePrivilegeIdForCustomerInBase(employee_privilege_id,employee_type)
}
