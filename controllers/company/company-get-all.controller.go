package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_company "bannayuu-web-admin/model/company"
	"bannayuu-web-admin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCompanyAll(c *gin.Context) {
	var companyResponseDb []model_company.CompanyGetAllResponse
	query := `select company_id,company_code
	,company_name,company_promotion
	,case when delete_flag = 'Y' then 'DISABLE'
	when current_timestamp < company_start_date then 'NOTOPEN'
	when current_timestamp > company_expire_date then 'EXPIRE'
	else 'NORMAL' end as status
	from m_company 
	
	order by company_code;`
	rows, err := db.GetDB().Raw(query).Rows()

	if err != nil {
		fmt.Printf("Get company error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), nil, fmt.Sprintf("Get company failed : %s", err))
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
