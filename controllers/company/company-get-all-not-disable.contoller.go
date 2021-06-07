package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_company "bannayuu-web-admin/model/company"
	"bannayuu-web-admin/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCompanyAllNotDisable(c *gin.Context) {
	var companyRequest model_company.CompanyGetAllRequest
	var companyResponseDb []model_company.CompanyGetAllResponse
	if err := c.ShouldBind(&companyRequest); err != nil {
		fmt.Printf("Combine Error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		return
	}
	query := `select company_id,company_code
	,company_name,company_promotion
	,case when delete_flag = 'Y' then 'DISABLE'
	when current_timestamp < company_start_date then 'NOTOPEN'
	when current_timestamp > company_expire_date then 'EXPIRE'
	else 'NORMAL' end as status
	from m_company where delete_flag = 'N' and current_timestamp < company_expire_date`
	if companyRequest.Company_code != ""  {
		query += ` and company_code = @company_code `
	}else if companyRequest.Company_name != ""{
		query += ` and company_name LIKE @company_name `
	}
	query += ` order by company_code;`
	likeStr := "%"
	company_name := fmt.Sprintf("%s%s%s", likeStr, companyRequest.Company_name, likeStr)
	rows, err := db.GetDB().Raw(query,
		sql.Named("company_code", companyRequest.Company_code),
		sql.Named("company_name", company_name)).Rows()

	fmt.Println(query)
	fmt.Println(companyRequest.Company_code)
	fmt.Println(companyRequest.Company_name)
	if err != nil {
		fmt.Printf("Get company not disable error : %s", err)
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

		fmt.Printf("Get company not disable successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": companyResponseDb, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), nil, "Get company successfully.")
	}
}
