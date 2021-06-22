package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_company "bannayuu-web-admin/model/company"
	"bannayuu-web-admin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnableCompany(c *gin.Context) {
	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("companyModel : %v ", jwtemployeeid)
	var companyModelReq model_company.CompanyDisableModelRequest
	if c.ShouldBind(&companyModelReq) == nil {
		//----------Save image
		rootCurrentDate := fmt.Sprintf("Company/%s", utils.GetDirectoryDate())
		imageName := utils.EncodeImageImage("COM_DIS")
		rootImages := fmt.Sprintf("%s/%s", constants.RootImages, rootCurrentDate)
		//----------check location path
		utils.CheckDirectory(rootImages)
		fileName := fmt.Sprintf("%s/%s", rootImages, imageName)
		errsaveimg := c.SaveUploadedFile(companyModelReq.Image, fileName)
		if errsaveimg != nil {
			c.String(http.StatusInternalServerError, constants.MessageImageNotFound)
			utils.WriteLog(utils.GetErrorLogCompanyFile(), constants.MessageImageNotFound)
			return
		}
		//----------Query
		saveEnableCompanyQuery(c, companyModelReq, jwtemployeeid, rootCurrentDate, imageName)
	} else {
		c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageDataNotCompletely})
	}
}

func saveEnableCompanyQuery(
	c *gin.Context,
	companyModelReq model_company.CompanyDisableModelRequest,
	jwtemployeeid interface{},
	rootCurrentDate string,
	imageName string,
) {
	//---------Convert obj setupdata to json string
	err_all_obj, all_obj := convertStrucToJSONStringAllForEnaable(companyModelReq, jwtemployeeid, fmt.Sprintf("%s/%s", rootCurrentDate, imageName))
	if err_all_obj {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
		utils.WriteLog(utils.GetErrorLogCompanyFile(), constants.MessageCovertObjTOJSONFailed)
		return
	}
	query := fmt.Sprintf(`
		with discomtb as (update m_company set
			company_remark = @remark
			,delete_flag = 'N'
			,update_by = @update_by
			,update_date = current_timestamp
			where company_id = @company_id
		)
		insert into log_company(
				lc_code
				,lc_name
				,lc_data
				,lc_type
				,create_by
				,create_date
				,company_id
			)values(
				fun_generate_uuid('LC'||trim(to_char(%s,'000')),5)
				,'เปิดให้บริการใหม่อีกครั้ง'
				,@log_data
				,'RESERVE'
				,@update_by
				,current_timestamp
				,@company_id
			)
			`, companyModelReq.Company_id)

	values := map[string]interface{}{
		"company_id": companyModelReq.Company_id,
		"remark":     companyModelReq.Remark,
		"update_by":  fmt.Sprint(jwtemployeeid),
		"log_data":   all_obj}

	if err, message := db.SaveTransactionDB(query, values); err {
		fmt.Printf("Enable company error : %s", message)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetErrorLogCompanyFile(), values, fmt.Sprintf("Enable company failed : %s", message))
	} else {
		fmt.Printf("Enable company successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), values, "Enable company successfully.")
	}
}

func convertStrucToJSONStringAllForEnaable(companyModelReq model_company.CompanyDisableModelRequest, jwtemployeeid interface{}, fileName string) (bool, string) {
	req_map := map[string]interface{}{
		"company_id": companyModelReq.Company_id,
		"remark":     companyModelReq.Remark,
		"update_by":  jwtemployeeid,
		"image":      fileName}
	err, setup_data := utils.ConvertInterfaceToJSON(req_map)
	if err {
		return true, ""
	}
	return false, setup_data
}
