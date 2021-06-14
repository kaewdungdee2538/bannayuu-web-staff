package interceptor

import (
	constants "bannayuu-web-admin/constants"
	home_intercep "bannayuu-web-admin/interceptor/home"
	user_model "bannayuu-web-admin/model/user"
	format_utls "bannayuu-web-admin/utils"
	// "bytes"
	// "encoding/json"
	// "io/ioutil"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckChangeMainCompanyUserValidateValuesInterceptor(c *gin.Context) {
	var userModel user_model.UserChangeMainCompanyRequestModel
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
	isErr, msg := checkValuesUserChangeMainCompany(userModel)
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

func checkValuesUserChangeMainCompany(userModel user_model.UserChangeMainCompanyRequestModel) (bool, string) {
	old_company_id := strings.TrimSpace(userModel.Old_company_id)
	new_company_id := strings.TrimSpace(userModel.New_company_id)
	employee_id := strings.TrimSpace(userModel.Employee_id)
	remark := strings.TrimSpace(userModel.Remark)
	if len(old_company_id) == 0 {
		return true, constants.MessageOldCompanyIdNotFound
	} else if format_utls.IsNotStringNumber(old_company_id) {
		return true, constants.MessageOldCompanyIdNotNumber
	} else if len(new_company_id) == 0 {
		return true, constants.MessageNewCompanyIdNotFound
	} else if format_utls.IsNotStringNumber(new_company_id) {
		return true, constants.MessageNewCompanyIdNotNumber
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
	errOldComId, _ := home_intercep.CheckValueCompanyIdNotDisavle(old_company_id)
	if errOldComId {
		return true, constants.MessageOldCompanyNotInBase
	}
	errNewComId, _ := home_intercep.CheckValueCompanyIdNotDisavle(new_company_id)
	if errNewComId {
		return true, constants.MessageNewCompanyNotInBase
	}
	errUserId, msgUserId := checkEmployeeIdInBase(employee_id, old_company_id)
	if errUserId {
		return true, msgUserId
	}
	return false, ""
}
