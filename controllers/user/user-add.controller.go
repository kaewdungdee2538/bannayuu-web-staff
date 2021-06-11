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

func AddUser(c *gin.Context) {
	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("jwt_employee_id : %v \n", jwtemployeeid)
	var userModelReq user_model.UserAddRequestModel

	if err := c.ShouldBind(&userModelReq); err == nil {
		fmt.Printf("user add Model : %v \n", userModelReq)
		//----------Insert User
		if err, message := insertUserQuery(&userModelReq, jwtemployeeid); err {
			utils.WriteLogInterface(utils.GetErrorLogUserFile(), nil, message)
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": message})
			return
		}
		//----------When Not Error
		fmt.Printf("add User successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogUserFile(), nil, "Add User successfully.")
	} else {
		utils.WriteLog(utils.GetAccessLogUserFile(), constants.MessageCombineFailed)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	}
}

func insertUserQuery(userModelReq *user_model.UserAddRequestModel, jwtemployeeid interface{}) (bool, string) {
	err_company_list_obj, company_list_obj := convertArrayToJson(userModelReq.Company_list)
	if err_company_list_obj {
		utils.WriteLog(utils.GetAccessLogCompanyFile(), constants.MessageCovertObjTOJSONFailed)
		return true, constants.MessageCovertObjTOJSONFailed
	}
	// ---------Convert obj to json string
	userInfo, err := json.Marshal(userModelReq)
	if err != nil {
		utils.WriteLogInterface(utils.GetErrorLogUserFile(), nil, constants.MessageCovertObjTOJSONFailed)
		return true, constants.MessageCovertObjTOJSONFailed
	}
	privilege_id := userModelReq.Employee_privilege_id
	company_id := userModelReq.Company_id
	
	if userModelReq.Employee_type == constants.EmployeeTypeOfUser{
		privilege_id = constants.PrivilegeOfUserTypeId
	}
	if userModelReq.Employee_type == constants.EmployeeTypeOfManagement{
		company_id = constants.CitCompany
	}
	query := `with insertemp as (
		insert into m_employee (
		employee_code,
		first_name_th,last_name_th,
		address,employee_mobile,
		employee_line,employee_email,
		username,passcode,
		employee_privilege_id,
		employee_status,
		company_id,
		company_list,
		create_by,create_date
	) values(
		fun_generate_uuid('EM',6),
		@first_name,@last_name,
		@address,@mobile,
		@line,@email,
		@username,crypt(@password,gen_salt('bf')) ,
		@employee_privilege_id,
		@status,
		@company_id,
		@company_list,
		@create_by,current_timestamp
	) RETURNING employee_id as emp_id,company_id as comid)
	 insert into log_employee(
		lep_code
		,lep_name
		,lep_data
		,lep_type
		,create_by
		,create_date
		,company_id
	)values(
		fun_generate_uuid('LEM'||trim(to_char((select comid FROM insertemp),'000')),5)
		,'สร้าง USER ใหม่'
		,@log_data
		,'CREATE'
		,@create_by
		,current_timestamp
		,(select comid FROM insertemp)
	)
	`
	values := map[string]interface{}{
		"first_name":            userModelReq.First_name,
		"last_name":             userModelReq.Last_name,
		"address":               userModelReq.Address,
		"mobile":                userModelReq.Mobile,
		"line":                  userModelReq.Line,
		"email":                 userModelReq.Email,
		"username":              userModelReq.Username,
		"password":              userModelReq.Password,
		"employee_privilege_id": privilege_id,
		"status":                userModelReq.Status,
		"company_id":            company_id,
		"company_list":          company_list_obj,
		"create_by":             fmt.Sprint(jwtemployeeid),
		"log_data":              string(userInfo),
	}
	fmt.Println(values)
	if err, message := db.SaveTransactionDB(query, values); err {
		fmt.Printf("add company Failed")
		utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), values, fmt.Sprintf("Add company Failed : %s", message))
		return true, message
	}
	return false, ""
}

func convertArrayToJson(companyList []int) (bool, string) {
	if fmt.Sprint(companyList[0])  == constants.CitCompany{
		return false,fmt.Sprintf("[%s]",constants.CitCompany)
	}
	urlsJson, err := json.Marshal(companyList)
	if err != nil {
		return true, ""
	}
	return false, string(urlsJson)
}
