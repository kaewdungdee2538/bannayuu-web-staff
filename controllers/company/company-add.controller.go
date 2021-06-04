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
		imageName := utils.EncodeImageImage("COM_ADD")
		rootImages := fmt.Sprintf("%s/%s", constants.RootImages, rootCurrentDate)
		//----------check location path
		utils.CheckDirectory(rootImages)
		fileName := fmt.Sprintf("%s/%s", rootImages, imageName)
		errsaveimg := c.SaveUploadedFile(companyModelReq.Image, fileName)
		if errsaveimg != nil {
			c.String(http.StatusInternalServerError, constants.MessageImageNotFound)
			utils.WriteLog(utils.GetAccessLogCompanyFile(), constants.MessageImageNotFound)
			return
		}
		//----------Query
		saveAddCompanyQuery(c, companyModelReq, jwtemployeeid, rootCurrentDate, imageName)
		fmt.Printf("companyModel : %v ", companyModelReq)

	} else {
		utils.WriteLog(utils.GetAccessLogCompanyFile(), constants.MessageCombineFailed)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	}
}
func convertStrucToJSONStringSetupCalForAdd(companyModelReq model_company.CompanyAddModelRequest) (bool, string) {
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

func convertStrucToJSONStringSetupVisitorInForAdd(companyModelReq model_company.CompanyAddModelRequest) (bool, string) {
	//---------Convert obj to json string
	var Booking_verify string
	if len(companyModelReq.Booking_verify) == 0 {
		Booking_verify = "qr"
	} else {
		Booking_verify = companyModelReq.Booking_verify
	}
	var Visitor_verify string
	if len(companyModelReq.Visitor_verify) == 0 {
		Visitor_verify = "identitycard"
	} else {
		Visitor_verify = companyModelReq.Visitor_verify
	}
	setup_data_map := map[string]interface{}{
		"booking_verify": Booking_verify,
		"visitor_verify": Visitor_verify}
	err, setup_data := utils.ConvertInterfaceToJSON(setup_data_map)
	if err {
		return true, ""
	}
	return false, setup_data
}

func convertStrucToJSONStringSetupVisitorOutForAdd(companyModelReq model_company.CompanyAddModelRequest) (bool, string) {
	//---------Convert obj to json string
	setup_data_map := map[string]interface{}{
		"booking_estamp_verify": companyModelReq.Booking_estamp_verify,
		"visitor_estamp_verify": companyModelReq.Visitor_estamp_verify}
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
	//---------Calcualte setup object
	err_setup_cal, setup_data_calculate := convertStrucToJSONStringSetupCalForAdd(companyModelReq)
	if err_setup_cal {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
		utils.WriteLog(utils.GetAccessLogCompanyFile(), constants.MessageCovertObjTOJSONFailed)
		return
	}
	//--------Visitor In Security setup object
	err_setup_visitor_in, setup_data_visitor_in := convertStrucToJSONStringSetupVisitorInForAdd(companyModelReq)
	if err_setup_visitor_in {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
		utils.WriteLog(utils.GetAccessLogCompanyFile(), constants.MessageCovertObjTOJSONFailed)
		return
	}
	//--------Visitor Out Security setup object
	err_setup_visitor_out, setup_data_visitor_out := convertStrucToJSONStringSetupVisitorOutForAdd(companyModelReq)
	if err_setup_visitor_out {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
		utils.WriteLog(utils.GetAccessLogCompanyFile(), constants.MessageCovertObjTOJSONFailed)
		return
	}

	err_all_obj, all_obj := convertStrucToJSONStringAllForAdd(companyModelReq, jwtemployeeid, fmt.Sprintf("%s/%s", rootCurrentDate, imageName))
	if err_all_obj {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
		utils.WriteLog(utils.GetAccessLogCompanyFile(), constants.MessageCovertObjTOJSONFailed)
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
		insertcomsetupcaltb as
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
				,@setup_data_calculate
				,'calculate_enable คือ เปิดใช้งานระบบคิดเงินหรือไม่ และ discount_split_from_day คือ เลือกว่าจะให้ลดนาทีแยกตามวัน หรือไม่'
				,(select comid FROM comtb)
			)),
		insertcomsetupvsintb as
		(insert into m_setup (
			setup_code
			,setup_name_en
			,setup_name_th
			,ref_setup_id
			,setup_data
			,setup_remark
			,company_id
			) values(
				fun_generate_uuid('MST001'||trim(to_char((select comid FROM comtb),'000')),0)
				,'Sub Visitor In Security'
				,'ความปลอดภัยหลักของ Visitor ขาเข้า'
				,1
				,@setup_data_visitor_in
				,'Booking จะต้องใช้ qr แล้ว เสียบบัตรประชาชนในการยืนยันตนเข้าหมู่บ้าน Visitor ต้องเสียบบัตรประชาชนทุกครั้ง'
				,(select comid FROM comtb)
			)),
			insertcomsetupvsouttb as
		(insert into m_setup (
			setup_code
			,setup_name_en
			,setup_name_th
			,ref_setup_id
			,setup_data
			,setup_remark
			,company_id
			) values(
				fun_generate_uuid('MST002'||trim(to_char((select comid FROM comtb),'000')),0)
				,'Sub Visitor Out Security'
				,'ความปลอดภัยของ Visitor ขาออก'
				,3
				,@setup_data_visitor_out
				,'booking จะต้อง estamp ก่อนถึงออกได้ และ visitor จะต้อง estamp ก่อนถึงจะออกได้'
				,(select comid FROM comtb)
			)),
		insertlogtb as (insert into log_company(
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
			))
			select func_addvisitorslot(200,@company_code,(select comid FROM comtb),@company_code,1,'001');
			`

	values := map[string]interface{}{
		"company_code":           companyModelReq.Company_code,
		"company_name":           companyModelReq.Company_name,
		"company_promotion":      companyModelReq.Company_promotion,
		"company_start_date":     companyModelReq.Company_start_date,
		"company_expire_date":    companyModelReq.Company_expire_date,
		"create_by":              jwtemployeeid,
		"setup_data_calculate":   setup_data_calculate,
		"setup_data_visitor_in":  setup_data_visitor_in,
		"setup_data_visitor_out": setup_data_visitor_out,
		"log_data":               all_obj}

	if err, message := db.SaveTransactionDB(query, values); err {
		fmt.Printf("add company Failed")
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), values, fmt.Sprintf("Add company Failed : %s", message))
	} else {
		fmt.Printf("add company successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), values, "Add company successfully.")
	}
}
