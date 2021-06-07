package interceptor

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	format_utls "bannayuu-web-admin/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CompanyGetIdModelRequest struct {
	Company_id string `form:"company_id"`
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

func checkValueCompanyId(Company_id string) (bool, string) {
	if len(Company_id) == 0 {
		return true, constants.MessageCompanyIdNotFound
	} else if format_utls.IsNotStringNumber(Company_id) {
		return true, constants.MessageCompanyIdNotNumber
	}
	return checkValuesGetCompanyId(Company_id);
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


