package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_company "bannayuu-web-admin/model/company"
	"bannayuu-web-admin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "io/ioutil"
)

func AddCompany(c *gin.Context) {
	// buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	// jsonString := string(buf)
	// fmt.Println(jsonString)
	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("companyModel : %v ", jwtemployeeid)
	var companyModelReq model_company.CompanyAddModelRequest
	if err := c.ShouldBind(&companyModelReq); err == nil {
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
		saveAddCompanyQuery(c, companyModelReq, jwtemployeeid, rootCurrentDate, imageName)
		fmt.Printf("companyModel : %v ", companyModelReq)

	} else {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	}
}
func convertStrucToJSONStringSetupForAdd(companyModelReq model_company.CompanyAddModelRequest) (bool, string) {
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
func convertStrucToJSONStringAllForAdd(companyModelReq model_company.CompanyAddModelRequest, jwtemployeeid interface{}, fileName string) (bool, string) {
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

func saveAddCompanyQuery(
	c *gin.Context,
	companyModelReq model_company.CompanyAddModelRequest,
	jwtemployeeid interface{},
	rootCurrentDate string,
	imageName string,
) {
	//---------Convert obj setupdata to json string
	err_setup, setup_data := convertStrucToJSONStringSetupForAdd(companyModelReq)
	if err_setup {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
		return
	}
	err_all_obj, all_obj := convertStrucToJSONStringAllForAdd(companyModelReq, jwtemployeeid, fmt.Sprintf("%s/%s", rootCurrentDate, imageName))
	if err_all_obj {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
		return
	}
	query := `
		with comtb as (
		insert into m_company(
			company_code
			,company_name
			,company_promotion
			,company_start_date
			,company_expire_date
			,create_by
			,create_date
		) values(
			@company_code
			,@company_name
			,@company_promotion
			,@company_start_date
			,@company_expire_date
			,@create_by
			,current_timestamp
			) 
			RETURNING company_id as comid
		),
		insertcomtb as
		(insert into m_setup (
			setup_code
			,setup_name_en
			,setup_name_th
			,ref_setup_id
			,setup_data
			,setup_remark
			,company_id
			) values(
				fun_generate_uuid('MST004'||trim(to_char((select comid FROM comtb),'000')),0)
				,'Sub Calculate setting'
				,'ตังค่าเกี่ยวกับระบบคิดเงิน'
				,8
				,@setup_data
				,'calculate_enable คือ เปิดใช้งานระบบคิดเงินหรือไม่ และ discount_split_from_day คือ เลือกว่าจะให้ลดนาทีแยกตามวัน หรือไม่'
				,(select comid FROM comtb)
			))
		insert into log_company(
				lc_code
				,lc_name
				,lc_data
				,lc_type
				,create_by
				,create_date
				,company_id
			)values(
				fun_generate_uuid('LC'||trim(to_char((select comid FROM comtb),'000')),5)
				,'เพิ่มโครงการใหม่'
				,@log_data
				,'CREATE'
				,@create_by
				,current_timestamp
				,(select comid FROM comtb)
			)
			`

	values := map[string]interface{}{
		"company_code":        companyModelReq.Company_code,
		"company_name":        companyModelReq.Company_name,
		"company_promotion":   companyModelReq.Company_promotion,
		"company_start_date":  companyModelReq.Company_start_date,
		"company_expire_date": companyModelReq.Company_expire_date,
		"create_by":           jwtemployeeid,
		"setup_data":          setup_data,
		"log_data":            all_obj}

	if err, message := db.SaveTransactionDB(query, values); err {
		fmt.Printf("add company error : %s", message)
		err, value_json := utils.ConvertInterfaceToJSON(values)
		if err {
			fmt.Printf("convert obj to json error : %s", constants.MessageCovertObjTOJSONFailed)
		}
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": message})
		utils.WriteLog(utils.GetAccessLogCompanyFile(), fmt.Sprintf("Add company failed !\nRequest : %v\n", string(value_json)))
	} else {
		fmt.Printf("add company successfully")
		err, value_json := utils.ConvertInterfaceToJSON(values)
		if err {
			fmt.Printf("convert obj to json error : %s", constants.MessageCovertObjTOJSONFailed)
		}
		c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageSuccess})
		utils.WriteLog(utils.GetAccessLogCompanyFile(), fmt.Sprintf("Add company successfully.\nRequest : %v\n", string(value_json)))
	}
}
