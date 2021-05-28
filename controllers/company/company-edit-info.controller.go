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

func EditInfoCompany(c *gin.Context) {
	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("companyModel : %v ", jwtemployeeid)
	var companyModelReq model_company.CompanyEditModelRequest
	if c.ShouldBind(&companyModelReq) == nil {
		//----------Save image
		rootCurrentDate := fmt.Sprintf("Company/%s", utils.GetDirectoryDate())
		imageName := utils.EncodeImageImage("ACM")
		rootImages := fmt.Sprintf("%s/%s", constants.RootImages, rootCurrentDate)
		//----------check location path
		utils.CheckDirectory(rootImages)
		fileName := fmt.Sprintf("%s/%s", rootImages, imageName)
		errsaveimg := c.SaveUploadedFile(companyModelReq.Image, fileName)
		if errsaveimg != nil {
			c.String(http.StatusInternalServerError, constants.MessageImageNotFound)
			return
		}
		//----------Query
		saveEditCompanyQuery(c, companyModelReq, jwtemployeeid, rootCurrentDate, imageName)
	}
}

func convertStrucToJSONStringForSetupForEdit(companyModelReq model_company.CompanyEditModelRequest) (bool, string) {
	//---------Convert obj to json string
	setup_data_map := map[string]interface{}{
		"calculate_enable":           companyModelReq.Calculate_enable,
		"except_time_split_from_day": companyModelReq.Except_time_split_from_day,
		"price_of_cardloss":          companyModelReq.Price_of_cardloss}
	err, setup_data := utils.ConvertInterfaceToJSON(setup_data_map)
	if err {
		return true, ""
	}
	return false, setup_data
}
func convertStrucToJSONStringAllForEdit(companyModelReq model_company.CompanyEditModelRequest, jwtemployeeid interface{}, fileName string) (bool, string) {
	req_map := map[string]interface{}{
		"company_code":               companyModelReq.Company_code,
		"company_name":               companyModelReq.Company_name,
		"company_promotion":          companyModelReq.Company_promotion,
		"company_start_date":         companyModelReq.Company_start_date,
		"company_expire_date":        companyModelReq.Company_expire_date,
		"create_by":                  jwtemployeeid,
		"calculate_enable":           companyModelReq.Calculate_enable,
		"except_time_split_from_day": companyModelReq.Except_time_split_from_day,
		"price_of_cardloss":          companyModelReq.Price_of_cardloss,
		"image":                      fileName}
	err, setup_data := utils.ConvertInterfaceToJSON(req_map)
	if err {
		return true, ""
	}
	return false, setup_data
}
func saveEditCompanyQuery(
	c *gin.Context,
	companyModelReq model_company.CompanyEditModelRequest,
	jwtemployeeid interface{},
	rootCurrentDate string,
	imageName string,
) {
	//---------Convert obj setupdata to json string
	err_setup, setup_data := convertStrucToJSONStringForSetupForEdit(companyModelReq)
	if err_setup {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
		return
	}
	err_all_obj, all_obj := convertStrucToJSONStringAllForEdit(companyModelReq, jwtemployeeid, fmt.Sprintf("%s/%s", rootCurrentDate, imageName))
	if err_all_obj {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
		return
	}
	query := fmt.Sprintf(`
		with editcomtb as (update m_company set
			company_code = @company_code
			,company_name = @company_name
			,company_promotion = @company_promotion
			,company_start_date = @company_start_date
			,company_expire_date = @company_expire_date
			,update_by = @update_by
			,update_date = current_timestamp
			where company_id = @company_id
		),
		editsetup as (update m_setup set
			setup_data = @setup_data
			where ref_setup_id = 8 and company_id = @company_id
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
				,'แก้ไขข้อมูลโครงการ'
				,@log_data
				,'EDIT'
				,@update_by
				,current_timestamp
				,@company_id
			)
			`,companyModelReq.Company_id);

	values := map[string]interface{}{
		"company_id":          companyModelReq.Company_id,
		"company_code":        companyModelReq.Company_code,
		"company_name":        companyModelReq.Company_name,
		"company_promotion":   companyModelReq.Company_promotion,
		"company_start_date":  companyModelReq.Company_start_date,
		"company_expire_date": companyModelReq.Company_expire_date,
		"update_by":           jwtemployeeid,
		"setup_data":          setup_data,
		"log_data":            all_obj}

	if err, message := db.SaveTransactionDB(query, values); err {
		fmt.Printf("edit company error : %s", message)
		err, value_json := utils.ConvertInterfaceToJSON(values)
		if err {
			fmt.Printf("convert obj to json error : %s", constants.MessageCovertObjTOJSONFailed)
		}
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": message})
		utils.WriteLog(utils.GetAccessLogCompanyFile(), fmt.Sprintf("Edit company failed !\nRequest : %v\n", string(value_json)))
	} else {
		fmt.Printf("Edit company successfully")
		err, value_json := utils.ConvertInterfaceToJSON(values)
		if err {
			fmt.Printf("convert obj to json error : %s", constants.MessageCovertObjTOJSONFailed)
		}
		c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageSuccess})
		utils.WriteLog(utils.GetAccessLogCompanyFile(), fmt.Sprintf("Edit company successfully.\nRequest : %v\n", string(value_json)))
	}
}
