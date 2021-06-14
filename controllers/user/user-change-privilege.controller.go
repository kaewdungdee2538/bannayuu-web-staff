package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	user_model "bannayuu-web-admin/model/user"
	"bannayuu-web-admin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "io/ioutil"
)

func ChangePrivilegeUser(c *gin.Context) {
	// buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	// jsonString := string(buf)
	// fmt.Println(jsonString)
	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("jwt_employee_id : %v ", jwtemployeeid)
	var userPrivilegeModelReq user_model.UserChangePrivilegeRequestModel
	if err := c.ShouldBind(&userPrivilegeModelReq); err == nil {
		//----------Save image
		rootCurrentDate := fmt.Sprintf("User/%s", utils.GetDirectoryDate())
		imageName := utils.EncodeImageImage("USER_CHANGEPRIVILEGE")
		rootImages := fmt.Sprintf("%s/%s", constants.RootImages, rootCurrentDate)
		//----------check location path
		utils.CheckDirectory(rootImages)
		fileName := fmt.Sprintf("%s/%s", rootImages, imageName)
		errsaveimg := c.SaveUploadedFile(userPrivilegeModelReq.Image, fileName)
		if errsaveimg != nil {
			c.String(http.StatusInternalServerError, constants.MessageImageNotFound)
			utils.WriteLog(utils.GetErrorLogUserFile(), constants.MessageImageNotFound)
			return
		}
		//----------Query
		saveChangePrivilegeUserQuery(c, userPrivilegeModelReq, jwtemployeeid, rootCurrentDate, imageName)
		fmt.Printf("userPrivilegeModelReq : %v ", userPrivilegeModelReq)

	} else {
		utils.WriteLog(utils.GetErrorLogUserFile(), constants.MessageCombineFailed)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	}
}

func convertStrucToJSONStringAllForChangePrivilegeUser(userPrivilegeModelReq user_model.UserChangePrivilegeRequestModel, jwtemployeeid interface{}, fileName string) (bool, string) {
	req_map := map[string]interface{}{
		"company_id":              userPrivilegeModelReq.Company_id,
		"employee_id":             userPrivilegeModelReq.Employee_id,
		"employee_privilege_id":   userPrivilegeModelReq.Employee_privilege_id,
		"employee_privilege_type": userPrivilegeModelReq.Employee_type,
		"update_by":               jwtemployeeid,
		"remark":                  userPrivilegeModelReq.Remark,
		"image":                   fileName}
	err, setup_data := utils.ConvertInterfaceToJSON(req_map)
	if err {
		return true, ""
	}
	return false, setup_data
}

func saveChangePrivilegeUserQuery(
	c *gin.Context,
	userPrivilegeModelReq user_model.UserChangePrivilegeRequestModel,
	jwtemployeeid interface{},
	rootCurrentDate string,
	imageName string,
) {
	//---------Convert obj setupdata to json string

	err_all_obj, all_obj := convertStrucToJSONStringAllForChangePrivilegeUser(userPrivilegeModelReq, jwtemployeeid, fmt.Sprintf("%s/%s", rootCurrentDate, imageName))
	if err_all_obj {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
		utils.WriteLog(utils.GetErrorLogUserFile(), constants.MessageCovertObjTOJSONFailed)
		return
	}
	query := `with updateemp as (
		update m_employee set
		employee_privilege_id = @employee_privilege_id,
		remark = @remark,
		update_by = @update_by,update_date = current_timestamp
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
		,'เปลี่ยนสิทธิ์การเข้าใช้งานของ USER'
		,@log_data
		,'CHANGEPRIVILEGE'
		,@update_by
		,current_timestamp
		,(select comid FROM updateemp)
	)
	`

	values := map[string]interface{}{
		"employee_privilege_id": userPrivilegeModelReq.Employee_privilege_id,
		"employee_id":           userPrivilegeModelReq.Employee_id,
		"company_id":            userPrivilegeModelReq.Company_id,
		"remark":                userPrivilegeModelReq.Remark,
		"update_by":             fmt.Sprint(jwtemployeeid),
		"log_data":              all_obj}

	if err, message := db.SaveTransactionDB(query, values); err {
		fmt.Printf("change privilege user Failed")
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetErrorLogUserFile(), values, fmt.Sprintf("Change privilege user Failed : %s", message))
	} else {
		fmt.Printf("change privilege user successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogUserFile(), values, "Change privilege user successfully.")
	}
}
