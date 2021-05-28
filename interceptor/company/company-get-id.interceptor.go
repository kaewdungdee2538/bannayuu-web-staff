package interceptor

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	format_utls "bannayuu-web-admin/utils"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type CompanyGetIdModelRequest struct {
	Company_id string
}
type CompanyGetIdModelResponse struct {
	Company_id int `json:"company_id"`
}

func GetIdCompanyValidateValuesInterceptor(c *gin.Context) {
	buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)
	var companyModel CompanyGetIdModelRequest
	err := json.Unmarshal([]byte(jsonString), &companyModel)
	if err != nil {
		//--------create error log
		format_utls.WriteLog(format_utls.GetErrorLogFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
		fmt.Printf("Error parsing JSON string - %s", err)
	}
	Company_id := companyModel.Company_id
	if err, msg := checkValueCompanyId(Company_id); err {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else if isErr, msg := checkValuesGetCompanyId(Company_id); isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		//-----forward request body middleware to endpoint
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr2
		c.Next()
	}
}

func checkValueCompanyId(Company_id string) (bool, string) {
	if len(Company_id) == 0 {
		return true, constants.MessageCompanyIdNotFound
	} else if format_utls.IsNotStringNumber(Company_id) {
		return true, constants.MessageCompanyIdNotNumber
	}
	return false, ""
}

func checkValuesGetCompanyId(comId string) (bool, string) {
	var companyIdObj CompanyGetIdModelResponse
	company_id := comId
	query := `select company_id from m_company
	where company_id = @company_id
	`
	rows, _ := db.GetDB().Raw(query, sql.Named("company_id", company_id)).Rows()
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
