package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_user "bannayuu-web-admin/model/user"
	"bannayuu-web-admin/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHomeInfo(c *gin.Context) {
	var userRequest model_user.UserGetByIdRequestModel
	var userResponse model_user.UserInfoGetResponseModel
	
	if err := c.ShouldBind(&userRequest); err != nil {
		fmt.Printf("Combine Error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		utils.WriteLogInterface(utils.GetErrorLogUserFile(), nil, fmt.Sprintf("Get Employee Info By Employee ID failed : %s", err))
		return
	}
	query := `select me.employee_id,me.employee_code,me.first_name_th,me.last_name_th,
	me.address,me.employee_telephone,me.employee_mobile,
	me.employee_line,me.employee_email,me.username,me.remark,
	me.employee_privilege_id,employee_privilege_name_th,employee_privilege_type,
	concat(mecreate.first_name_th,' ',mecreate.last_name_th) as create_by,
	to_char(me.create_date,'YYYY-MM-DD HH24:MI:SS') as create_date,
	concat(meupdate.first_name_th,' ',meupdate.last_name_th) as update_by,
	to_char(me.update_date,'YYYY-MM-DD HH24:MI:SS') as update_date,
	concat(medelete.first_name_th,' ',medelete.last_name_th) as delete_by,
	to_char(me.delete_date,'YYYY-MM-DD HH24:MI:SS') as delete_date,
	case when me.delete_flag = 'Y' then 'DISABLE' else 'NORMAL' end delete_flag
	from m_employee me
	left join m_employee_privilege mep
	on me.employee_privilege_id = mep.employee_privilege_id
	left join m_employee mecreate
	on me.create_by = mecreate.employee_id
	left join m_employee meupdate
	on me.update_by = meupdate.employee_id
	left join m_employee medelete
	on me.delete_by = medelete.employee_id
	where me.company_id = @company_id and me.employee_id = @employee_id
	limit 1`

	rows, err := db.GetDB().Raw(query,
		sql.Named("company_id", userRequest.Company_id),
		sql.Named("employee_id", userRequest.Employee_id),
		).Rows()

	if err != nil {
		fmt.Printf("Get Employee Info By Employee ID error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetErrorLogUserFile(), nil, fmt.Sprintf("Get Employee Info By Employee ID failed : %s", err))
		defer rows.Close()
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&userResponse)
			db.GetDB().ScanRows(rows, &userResponse)
			// do something
		}

		fmt.Printf("Get Employee Info By Employee ID successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": userResponse, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogUserFile(), nil, "Get Employee Info By Employee ID successfully.")
	}
}
