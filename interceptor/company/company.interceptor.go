package interceptor

import (
	constants "bannayuu-web-admin/constants"
	format_utls "bannayuu-web-admin/utils"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"io/ioutil"
	"net/http"
	"strings"
)

type CompanyAddModelRequest struct {
	Company_code        string
	Company_name        string
	Company_promotion   string
	Company_start_date  string
	Company_expire_date string
	Create_by           int
}

func AddCompanyValidateValues(c *gin.Context) {
	buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)
	var companyModel CompanyAddModelRequest
	err := json.Unmarshal([]byte(jsonString), &companyModel)
	if err != nil {
		//--------create error log
		constants.WriteLog(constants.GetErrorLogFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
		fmt.Printf("Error parsing JSON string - %s", err)
	}
	isErr, msg := checkValues(companyModel)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		//-----forward request body middleware to endpoint
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr2
		c.Next()
	}
}

func checkValues(companyModel CompanyAddModelRequest) (bool, string) {
	Company_code := strings.TrimSpace(companyModel.Company_code)
	Company_name := strings.TrimSpace(companyModel.Company_name)
	Company_promotion := strings.TrimSpace(companyModel.Company_promotion)
	Company_start_date := strings.TrimSpace(companyModel.Company_start_date)
	Company_expire_date := strings.TrimSpace(companyModel.Company_expire_date)
	if len(Company_code) == 0 {
		return true, constants.MessageCompanyCodeNotFount
	} else if format_utls.IsNotStringEngOtNumber(Company_code) {
		return true, constants.MessageCompanyCodeIsSpecialProhibit
	} else if len(Company_name) == 0 {
		return true, constants.MessageCompanyNameNotFount
	} else if format_utls.IsNotStringAlphabet(Company_name) {
		return true, constants.MessageCompanyNameIsSpecialProhibit
	} else if len(Company_promotion) == 0 {
		return true, constants.MessageCompanyProNotFound
	} else if format_utls.IsNotStringEngOtNumber(Company_promotion) {
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
