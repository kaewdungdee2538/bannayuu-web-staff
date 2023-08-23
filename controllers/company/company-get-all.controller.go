package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_company "bannayuu-web-admin/model/company"
	"bannayuu-web-admin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCompanyAll(c *gin.Context) {
	var companyRequest model_company.CompanyGetAllRequest
	var companyResponseDb []model_company.CompanyGetAllResponse
	if err := c.ShouldBind(&companyRequest); err != nil {
		fmt.Printf("Combine Error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		return
	}
	query := `with input_data AS (
		SELECT
			$1::TEXT AS company_code_or_name
	)
	select company_id,company_code
	,company_name,company_promotion
	,case when delete_flag = 'Y' then 'DISABLE'
	when current_timestamp < company_start_date then 'NOTOPEN'
	when current_timestamp > company_expire_date then 'EXPIRE'
	else 'NORMAL' end as status
	from m_company where company_id IS NOT NULL`
	if companyRequest.Company_code_or_name != "" {
		query += ` and (company_code LIKE (SELECT company_code_or_name FROM input_data) or company_name LIKE (SELECT company_code_or_name FROM input_data))`
	}
	query += ` order by company_code;`
	likeStr := "%"
	Company_code_name := fmt.Sprintf("%s%s%s", likeStr, companyRequest.Company_code_or_name, likeStr)

	rows, err := db.GetDB().Raw(query,
		Company_code_name,
	).Rows()

	if err != nil {
		fmt.Printf("Get company error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetErrorLogCompanyFile(), nil, fmt.Sprintf("Get company failed : %s", err))
		defer rows.Close()
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&companyResponseDb)
			db.GetDB().ScanRows(rows, &companyResponseDb)
			// do something
		}

		fmt.Printf("Get company successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": companyResponseDb, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), nil, "Get company successfully.")
	}
}
