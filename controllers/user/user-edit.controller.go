package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	user_model "bannayuu-web-admin/model/user"
	"bannayuu-web-admin/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditUser(c *gin.Context) {
	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("jwt_employee_id : %v \n", jwtemployeeid)
	var userModelReq user_model.UserEditInfoRequestModel

	if err := c.ShouldBind(&userModelReq); err == nil {
		fmt.Printf("user edit Model : %v \n", userModelReq)
		//----------Insert User
		if err, message := updateUserQuery(&userModelReq, jwtemployeeid); err {
			utils.WriteLogInterface(utils.GetErrorLogUserFile(), nil, message)
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": message})
			return
		}
		//----------When Not Error
		fmt.Printf("edit User successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogUserFile(), nil, "Edit User successfully.")
	} else {
		utils.WriteLog(utils.GetAccessLogUserFile(), constants.MessageCombineFailed)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	}
}

func updateUserQuery(userModelReq *user_model.UserEditInfoRequestModel, jwtemployeeid interface{}) (bool, string) {

	// ---------Convert obj to json string
	userInfo, err := json.Marshal(userModelReq)
	if err != nil {
		utils.WriteLogInterface(utils.GetErrorLogUserFile(), nil, constants.MessageCovertObjTOJSONFailed)
		return true, constants.MessageCovertObjTOJSONFailed
	}
	query := `with updateemp as (
		update m_employee set
		first_name_th = @first_name,last_name_th = @last_name,
		address = @address,employee_mobile = @mobile,employee_line = @line,
		employee_email = @email,update_by = @update_by,update_date = current_timestamp
		where employee_id = @employee_id and company_id = @company_id
	 RETURNING employee_id as emp_id,company_id as comid)
	 insert into log_employee(
		lep_code
		,lep_name
		,lep_data
		,lep_type
		,create_by
		,create_date
		,company_id
	)values(
		fun_generate_uuid('LEM'||trim(to_char((select comid FROM updateemp),'000')),5)
		,'แก้ไขข้อมูล USER'
		,@log_data
		,'UPDATEINFO'
		,@update_by
		,current_timestamp
		,(select comid FROM updateemp)
	)
	`
	values := map[string]interface{}{
		"first_name": userModelReq.First_name,
		"last_name":  userModelReq.Last_name,
		"address":    userModelReq.Address,
		"mobile":     userModelReq.Mobile,
		"line":       userModelReq.Line,
		"email":      userModelReq.Email,
		"update_by":  fmt.Sprint(jwtemployeeid),
		"employee_id": userModelReq.Employee_id,
		"company_id" : userModelReq.Company_id,
		"log_data":   string(userInfo),
	}
	fmt.Println(values)
	if err, message := db.SaveTransactionDB(query, values); err {
		fmt.Printf("edit company Failed")
		utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), values, fmt.Sprintf("Edit company Failed : %s", message))
		return true, message
	}
	return false, ""
}
