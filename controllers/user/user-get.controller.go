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

func GetUserAll(c *gin.Context) {
	var userRequest model_user.UserGetRequestModel
	var userResponse []model_user.UserGetResponseModel
	
	if err := c.ShouldBind(&userRequest); err != nil {
		fmt.Printf("Combine Error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		utils.WriteLogInterface(utils.GetErrorLogUserFile(), nil, fmt.Sprintf("Get Employee Info By Employee ID failed : %s", err))
		return
	}
	query := `select 
	employee_id,employee_code,
	first_name_th,last_name_th,
	username,
	employee_privilege_type
	from m_employee me
	left join m_employee_privilege mep
	on me.employee_privilege_id = mep.employee_privilege_id
	where me.company_id = @company_id
	and (username LIKE @full_name or first_name_th LIKE @full_name or last_name_th LIKE @full_name)
	order by first_name_th,last_name_th,username;`

	likeStr := "%"
	full_name := fmt.Sprintf("%s%s%s", likeStr, userRequest.Full_name, likeStr)

	rows, err := db.GetDB().Raw(query,
		sql.Named("company_id", userRequest.Company_id),
		sql.Named("full_name", full_name),
		).Rows()

	if err != nil {
		fmt.Printf("Get Employee All error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetErrorLogUserFile(), nil, fmt.Sprintf("Get Employee All failed : %s", err))
		defer rows.Close()
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&userResponse)
			db.GetDB().ScanRows(rows, &userResponse)
			// do something
		}

		fmt.Printf("Get Employee Info All successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": userResponse, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogUserFile(), nil, "Get Employee Info All successfully.")
	}
}
