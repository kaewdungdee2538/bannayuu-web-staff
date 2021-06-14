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

func GetPrivilege(c *gin.Context) {
	var userResponse []model_user.UserGetPrivilegeResponseModel

	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("jwt_employee_id : %v \n", jwtemployeeid)

	query := `	with emptbl as (select me.employee_privilege_id as privilege_id from m_employee me
		left join m_employee_privilege mep
		on me.employee_privilege_id = mep.employee_privilege_id
		where me.employee_id = @employee_id
		)
		select employee_privilege_id,
		employee_privilege_name_th,
		employee_privilege_name_en,
		employee_privilege_type
		from m_employee_privilege
		where employee_privilege_id not in (7,(select privilege_id from emptbl))
		order by employee_privilege_id`

	rows, err := db.GetDB().Raw(query,
		sql.Named("employee_id", jwtemployeeid),
	).Rows()

	if err != nil {
		fmt.Printf("Get Employee privilege error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetErrorLogUserFile(), nil, fmt.Sprintf("Get Employee privilege failed : %s", err))
		defer rows.Close()
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&userResponse)
			db.GetDB().ScanRows(rows, &userResponse)
			// do something
		}

		fmt.Printf("Get Employee privilege successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": userResponse, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogUserFile(), nil, "Get Employee privilege successfully.")
	}
}
