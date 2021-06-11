package interceptor

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_company "bannayuu-web-admin/model/company"
	format_utls "bannayuu-web-admin/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"strconv"
)

func EditCompanyValidateValuesInterceptor(c *gin.Context) {
	var companyModel model_company.CompanyEditModelRequest
	if err := c.ShouldBind(&companyModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		c.Abort()
		return
	}
	fmt.Print(companyModel)
	isErr, msg := checkValuesEditCompany(companyModel)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		c.Next()
	}
}

func checkValuesEditCompany(companyModel model_company.CompanyEditModelRequest) (bool, string) {
	Company_id := strings.TrimSpace(companyModel.Company_id)
	Company_code := strings.TrimSpace(companyModel.Company_code)
	Company_name := strings.TrimSpace(companyModel.Company_name)
	Company_promotion := strings.TrimSpace(companyModel.Company_promotion)
	Company_start_date := strings.TrimSpace(companyModel.Company_start_date)
	Company_expire_date := strings.TrimSpace(companyModel.Company_expire_date)
	remark := strings.TrimSpace(companyModel.Remark)
	if len(Company_id) == 0 {
		return true, constants.MessageCompanyIdNotFound
	} else if format_utls.IsNotStringNumber(Company_id) {
		return true, constants.MessageCompanyIdNotNumber
	} else if len(Company_code) == 0 {
		return true, constants.MessageCompanyCodeNotFount
	} else if format_utls.IsNotStringEngOrNumber(Company_code) {
		return true, constants.MessageCompanyCodeIsSpecialProhibit
	} else if len(Company_name) == 0 {
		return true, constants.MessageCompanyNameNotFount
	} else if format_utls.IsNotStringAlphabetRemark(Company_name) {
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
	} else if len(remark) == 0 {
		return true, constants.MessageRemarkNotFount
	} else if len(remark) < 10 {
		return true, constants.MessageRemarkIsLower10Character
	} else if format_utls.IsNotStringAlphabetRemark(remark) {
		return true, constants.MessageRemarkProhibitSpecial
	} else if err, msg := checkValuesGetCompanyId(Company_id); err {
		return true, msg
	}
	return checkCompanyDuplicateWhenEdit(Company_id, Company_code, Company_name)
}

func checkCompanyDuplicateWhenEdit(comId string, comCode string, comName string) (bool, string) {
	var companyIdObj CompanyGetIdModelResponse
	company_code := comCode
	company_name := comName
	company_id,_ := strconv.ParseInt(comId,10,64); 
	query := `select company_id from m_company
	where delete_flag = 'N' 
	and company_id != @company_id
	and (company_code = @company_code
		or company_name = @company_name)
	;
	`
	fmt.Println(query)
	fmt.Println(company_id)
	fmt.Println(company_code)
	fmt.Println(company_name)
	rows, err := db.GetDB().Raw(query, sql.Named("company_id", company_id),sql.Named("company_code", company_code),sql.Named("company_name", company_name)).Rows();
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&companyIdObj)
		db.GetDB().ScanRows(rows, &companyIdObj)
		// do something
	}
	if companyIdObj.Company_id > 0 {
		return true, constants.MessageCompanyIsDuplicateInBase
	}
	return false, ""
}
