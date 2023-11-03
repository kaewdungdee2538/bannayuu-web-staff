package interceptor

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	format_utls "bannayuu-web-admin/utils"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyGetIdModelRequest struct {
	Company_id string `form:"company_id"`
}

type CompanyGetIdJSONModelRequest struct {
	Company_id int `json:"company_id"`
}
type CompanyGetIdModelResponse struct {
	Company_id int `json:"company_id"`
}

func GetIdCompanyValidateValuesInterceptor(c *gin.Context) {
	var companyModel CompanyGetIdModelRequest
	if err := c.ShouldBind(&companyModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		c.Abort()
		return
	}
	fmt.Print(companyModel)
	isErr, msg := checkValueCompanyId(companyModel.Company_id)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		c.Next()
	}
}

func GetIdCompanyJsonValidateValuesInterceptor(c *gin.Context) {
	var request CompanyGetIdJSONModelRequest
	buf, _ := io.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)

	err := json.Unmarshal([]byte(jsonString), &request)

	if err != nil {
		//--------create error log
		format_utls.WriteLog(format_utls.GetErrorLogUserFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
		fmt.Printf("Error parsing JSON string - %s", err)
	}
	isErr, msg := checkValueCompanyId(fmt.Sprint(request.Company_id))
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		// ---------Convert obj to json string
		userInfo, err := json.Marshal(request)
		if err != nil {
			format_utls.WriteLogInterface(format_utls.GetErrorLogUserFile(), nil, constants.MessageCovertObjTOJSONFailed)
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
			return
		}
		// -----forward request body middleware to endpoint
		rdr2 := io.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf("%v", string(userInfo)))))
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
	return checkValuesGetCompanyId(Company_id)
}

func checkValuesGetCompanyId(comId string) (bool, string) {
	var companyIdObj CompanyGetIdModelResponse
	company_id := comId
	query := `select company_id from m_company
	where company_id = @company_id;
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
