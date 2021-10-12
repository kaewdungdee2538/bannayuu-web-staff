package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	user_reset_password_model "bannayuu-web-admin/model/user"
	"bannayuu-web-admin/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResetPasswordUser(c *gin.Context) {
	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("jwt_employee_id : %v \n", jwtemployeeid)
	var userModelReq user_reset_password_model.UserResetPasswordRequestModel

	if err := c.ShouldBind(&userModelReq); err == nil {
		fmt.Printf("user reset password Model : %v \n", userModelReq)
		//----------Insert User
		if err, message := genResetPasswordUserQuery(&userModelReq, jwtemployeeid); err {
			utils.WriteLogInterface(utils.GetErrorLogUserResetPasswordFile(), nil, message)
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": message})
			return
		} else {
			//----------When Not Error
			fmt.Printf("user reset successfully")
			c.JSON(http.StatusOK, gin.H{"error": false, "result": message, "message": message})
			utils.WriteLogInterface(utils.GetAccessLogUserResetPasswordFile(), nil, "user reset successfully.")
		}
	} else {
		utils.WriteLog(utils.GetErrorLogUserResetPasswordFile(), constants.MessageCombineFailed)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	}
}

func genResetPasswordUserQuery(userModelReq *user_reset_password_model.UserResetPasswordRequestModel, jwtemployeeid interface{}) (bool, string) {
	var userResponse user_reset_password_model.UserResetPasswordResponseModel
	query := `with insertreset as (
		insert into t_employee_password_info(
			tepi_code,tepi_status
			,employee_id,tepi_hold
			,remark
			,create_by,create_date
			,company_id
		) values(
			fun_generate_uuid('PRE',8),'GEN_RESET'
			,@employee_id,@hold_time
			,@remark
			,@create_by,current_timestamp
			,@company_id
		)
		RETURNING tepi_id,tepi_code,employee_id,company_id
	)
	select tepi_id,tepi_code,employee_id,company_id from insertreset;
	`
	values := map[string]interface{}{
		"employee_id": userModelReq.Employee_id,
		"hold_time":   userModelReq.Hold_time,
		"remark":      userModelReq.Remark,
		"create_by":   fmt.Sprint(jwtemployeeid),
		"company_id":  userModelReq.Company_id,
	}
	fmt.Println(values)
	rows, err := db.GetDB().Raw(query,
		sql.Named("employee_id", userModelReq.Employee_id),
		sql.Named("hold_time", userModelReq.Hold_time),
		sql.Named("remark", userModelReq.Remark),
		sql.Named("create_by", fmt.Sprint(jwtemployeeid)),
		sql.Named("company_id", userModelReq.Company_id)).Rows()

	if err != nil {
		errMsg := fmt.Sprintf("user reset error : %s", err)
		utils.WriteLogInterface(utils.GetErrorLogUserResetPasswordFile(), nil, errMsg)
		defer rows.Close()
		return true, errMsg
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&userResponse)
			db.GetDB().ScanRows(rows, &userResponse)
			// do something
		}
	}
	fmt.Println(userResponse)
	uri_response := fmt.Sprintf("%s/%d/%d/%s",constants.WEB_MANAGEMENT_RESET_USER,userResponse.Company_id, userResponse.Employee_id, userResponse.Tepi_code)
	return false, uri_response
}
