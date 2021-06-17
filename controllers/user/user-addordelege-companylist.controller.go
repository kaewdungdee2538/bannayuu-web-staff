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

func AddOrDeleteCompanyListUser(c *gin.Context) {
	// buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	// jsonString := string(buf)
	// fmt.Println(jsonString)
	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("jwt_employee_id : %v ", jwtemployeeid)
	var userCompanyModelReq user_model.UserAddOrDeleteCompanyListRequestModel
	if err := c.ShouldBind(&userCompanyModelReq); err == nil {
		//----------Save image
		rootCurrentDate := fmt.Sprintf("User/%s", utils.GetDirectoryDate())
		imageName := utils.EncodeImageImage("USER_ADDORDELETECOMPANYLIST")
		rootImages := fmt.Sprintf("%s/%s", constants.RootImages, rootCurrentDate)
		//----------check location path
		utils.CheckDirectory(rootImages)
		fileName := fmt.Sprintf("%s/%s", rootImages, imageName)
		errsaveimg := c.SaveUploadedFile(userCompanyModelReq.Image, fileName)
		if errsaveimg != nil {
			c.String(http.StatusInternalServerError, constants.MessageImageNotFound)
			utils.WriteLog(utils.GetErrorLogUserFile(), constants.MessageImageNotFound)
			return
		}
		//----------Query
		saveAddOrDeleteCompanyListUserQuery(c, userCompanyModelReq, jwtemployeeid, rootCurrentDate, imageName)
		fmt.Printf("userCompanyListModelReq : %v ", userCompanyModelReq)

	} else {
		utils.WriteLog(utils.GetErrorLogUserFile(), constants.MessageCombineFailed)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	}
}

func convertStrucToJSONStringAllForAddOrDeleteCompanyListUser(userCompanyListModelReq user_model.UserAddOrDeleteCompanyListRequestModel, jwtemployeeid interface{}, fileName string) (bool, string) {
	req_map := map[string]interface{}{
		"company_id":   userCompanyListModelReq.Company_id,
		"company_list": userCompanyListModelReq.Company_list,
		"employee_id":  userCompanyListModelReq.Employee_id,
		"update_by":    jwtemployeeid,
		"remark":       userCompanyListModelReq.Remark,
		"image":        fileName}
	err, setup_data := utils.ConvertInterfaceToJSON(req_map)
	if err {
		return true, ""
	}
	return false, setup_data
}

func saveAddOrDeleteCompanyListUserQuery(
	c *gin.Context,
	userPrivilegeModelReq user_model.UserAddOrDeleteCompanyListRequestModel,
	jwtemployeeid interface{},
	rootCurrentDate string,
	imageName string,
) {
	// //---------Convert obj setupdata to json string
	// err_company_list_obj, company_list_obj := convertCompanyArrayToJson(userPrivilegeModelReq.Company_list)
	// if err_company_list_obj {
	// 	c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
	// 	utils.WriteLog(utils.GetAccessLogCompanyFile(), constants.MessageCovertObjTOJSONFailed)
	// 	return 
	// }

	err_all_obj, all_obj := convertStrucToJSONStringAllForAddOrDeleteCompanyListUser(userPrivilegeModelReq, jwtemployeeid, fmt.Sprintf("%s/%s", rootCurrentDate, imageName))
	if err_all_obj {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
		utils.WriteLog(utils.GetErrorLogUserFile(), constants.MessageCovertObjTOJSONFailed)
		return
	}
	query := `with updateemp as (
		update m_employee set
		company_list = @company_list,
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
		,'เพิ่มหรือลดโครงการที่ดูแล USER'
		,@log_data
		,'ADDORDELETECOMPANYLIST'
		,@update_by
		,current_timestamp
		,(select comid FROM updateemp)
	)
	`

	values := map[string]interface{}{
		"company_id": userPrivilegeModelReq.Company_id,
		"company_list": userPrivilegeModelReq.Company_list,
		"employee_id":    userPrivilegeModelReq.Employee_id,
		"remark":         userPrivilegeModelReq.Remark,
		"update_by":      fmt.Sprint(jwtemployeeid),
		"log_data":       all_obj}

	if err, message := db.SaveTransactionDB(query, values); err {
		fmt.Printf("add or delete Company list user Failed")
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetErrorLogUserFile(), values, fmt.Sprintf("Add or delete Company list user Failed : %s", message))
	} else {
		fmt.Printf("Add or delete Company list user successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogUserFile(), values, "Add or delete Company list user successfully.")
	}
}


