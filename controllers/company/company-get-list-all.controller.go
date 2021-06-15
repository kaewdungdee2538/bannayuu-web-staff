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

func GetCompanyListAll(c *gin.Context) {
	var companyResponseDb []model_company.CompanyGetAllResponse
	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("jwt_employee_id : %v ", jwtemployeeid)
	
	query := `with emptbl as (select me.employee_privilege_id as privilege_id from m_employee me
		left join m_employee_privilege mep
		on me.employee_privilege_id = mep.employee_privilege_id
		where me.employee_id = @employee_id
		),
	comtbl as (select case when (select privilege_id from emptbl) <= 6 then 999 else 0 end as comid)
		select company_id,company_code
		,company_name,company_promotion
		,case when delete_flag = 'Y' then 'DISABLE'
		when current_timestamp < company_start_date then 'NOTOPEN'
		when current_timestamp > company_expire_date then 'EXPIRE'
		else 'NORMAL' end as status
		from m_company 
		where company_id  not in (select comid from comtbl)
		order by company_name`
	rows, err := db.GetDB().Raw(query,
		sql.Named("employee_id", jwtemployeeid),
		).Rows()

	if err != nil {
		fmt.Printf("Get company list error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetErrorLogCompanyFile(), nil, fmt.Sprintf("Get company List failed : %s", err))
		defer rows.Close()
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&companyResponseDb)
			db.GetDB().ScanRows(rows, &companyResponseDb)
			// do something
		}

		fmt.Printf("Get company List successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": companyResponseDb, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), nil, "Get company List successfully.")
	}
}
