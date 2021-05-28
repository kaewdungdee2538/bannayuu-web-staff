package interceptor

import (
	constants "bannayuu-web-admin/constants"
	model_company "bannayuu-web-admin/model/company"
	format_utls "bannayuu-web-admin/utils"
	// "bytes"
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddCompanyValidateValuesInterceptor(c *gin.Context) {
	var companyModel model_company.CompanyAddModelRequest
	if err := c.ShouldBind(&companyModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		c.Abort()
		return
	}
	fmt.Print(companyModel)
	// buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	// jsonString := string(buf)

	// err := json.Unmarshal([]byte(jsonString), &companyModel)
	// if err != nil {
	// 	//--------create error log
	// 	format_utls.WriteLog(format_utls.GetErrorLogFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
	// 	fmt.Printf("Error parsing JSON string - %s", err)
	// }
	isErr, msg := checkValuesAddCompany(companyModel)
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

func checkValuesAddCompany(companyModel model_company.CompanyAddModelRequest) (bool, string) {
	Company_code := strings.TrimSpace(companyModel.Company_code)
	Company_name := strings.TrimSpace(companyModel.Company_name)
	Company_promotion := strings.TrimSpace(companyModel.Company_promotion)
	Company_start_date := strings.TrimSpace(companyModel.Company_start_date)
	Company_expire_date := strings.TrimSpace(companyModel.Company_expire_date)
	if len(Company_code) == 0 {
		return true, constants.MessageCompanyCodeNotFount
	} else if format_utls.IsNotStringEngOrNumber(Company_code) {
		return true, constants.MessageCompanyCodeIsSpecialProhibit
	} else if len(Company_name) == 0 {
		return true, constants.MessageCompanyNameNotFount
	} else if format_utls.IsNotStringAlphabet(Company_name) {
		return true, constants.MessageCompanyNameIsSpecialProhibit
	} else if len(Company_promotion) == 0 {
		return true, constants.MessageCompanyProNotFound
	} else if format_utls.IsNotStringEngOrNumber(Company_promotion) {
		return true, constants.MessageCompanyProIsSpecialProhibit
	} else if len(Company_start_date) == 0 {
		return true, constants.MessageDateStartNotFound
	} else if format_utls.IsNotFormatTime(Company_start_date) {
		return true, constants.MessageDateStarFormatNotValid
	} else if len(Company_expire_date) == 0 {
		return true, constants.MessageDateStopNotFound
	} else if format_utls.IsNotFormatTime(Company_expire_date) {
		return true, constants.MessageDateStopFormatNotValid
	}
	return false, ""
}
