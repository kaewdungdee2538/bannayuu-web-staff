package interceptor

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	home_model "bannayuu-web-admin/model/home"
	company_model "bannayuu-web-admin/model/company"
	format_utls "bannayuu-web-admin/utils"
	"bytes"
	// "database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)


func CheckAddHomeArrayValuesInterceptor(c *gin.Context) {
	var companyModel home_model.HomeAddRequestModel
	buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)

	err := json.Unmarshal([]byte(jsonString), &companyModel)

	if err != nil {
		//--------create error log
		// format_utls.WriteLog(format_utls.GetErrorLogFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
		fmt.Printf("Error parsing JSON string - %s", err)
	}

	// if err := c.ShouldBind(&companyModel); err != nil {
	// 	c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	// 	c.Abort()
	// 	return
	// }
	fmt.Print(companyModel)
	isErr, msg := CheckValueCompanyIdNotDisavle(companyModel.Company_id)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		// ---------Convert obj to json string
		companyInfo, err := json.Marshal(companyModel)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
			return
		}
		// -----forward request body middleware to endpoint
		rdr2 := ioutil.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf("%v", string(companyInfo)))))
		c.Request.Body = rdr2
		c.Next()
	}
}

func CheckValueCompanyIdNotDisavle(Company_id string) (bool, string) {
	if len(Company_id) == 0 {
		return true, constants.MessageCompanyIdNotFound
	} else if format_utls.IsNotStringNumber(Company_id) {
		return true, constants.MessageCompanyIdNotNumber
	}
	return checkValuesGetCompanyIdNotDisable(Company_id)
}

func checkValuesGetCompanyIdNotDisable(comId string) (bool, string) {
	var companyIdObj company_model.CompanyGetIdNotDisableResponseModel
	company_id := comId
	query := fmt.Sprintf(`select company_id from m_company
	where delete_flag = 'N' and company_id = %v;`,company_id)
	rows, _ := db.GetDB().Raw(query).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&companyIdObj)
		db.GetDB().ScanRows(rows, &companyIdObj)
		// do something
	}
	if companyIdObj.Company_id == 0 {
		return true, constants.MessageCompanyNotInBase
	}
	return false, ""
}
