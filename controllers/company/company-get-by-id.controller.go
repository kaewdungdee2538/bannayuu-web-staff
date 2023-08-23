package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_company "bannayuu-web-admin/model/company"
	"bannayuu-web-admin/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCompanyById(c *gin.Context) {
	var companyRequestDb model_company.CompanyGetByIdResquest
	var companyResponseDb model_company.CompanyGetByIdFromDBResponse
	if err := c.ShouldBind(&companyRequestDb); err != nil {
		fmt.Printf("Combine Error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		return
	}
	query := `select mc.company_id,company_code
	,company_name,company_promotion
	,case when mc.delete_flag = 'Y' then 'DISABLE'
	when current_timestamp < company_start_date then 'NOTOPEN'
	when current_timestamp > company_expire_date then 'EXPIRE'
	else 'NORMAL' end as status
	,to_char(company_start_date,'YYYY-MM-DD HH24:MI:SS') as company_start_date
	,to_char(company_expire_date,'YYYY-MM-DD HH24:MI:SS') as company_expire_date
	,company_remark
	,(select concat(first_name_th,' ',last_name_th) from m_employee where employee_id = mc.create_by) as create_by
	,to_char(mc.create_date,'YYYY-MM-DD HH24:MI:SS') as create_date
	,(select concat(first_name_th,' ',last_name_th) from m_employee where employee_id = mc.update_by) as update_by
	,to_char(mc.update_date,'YYYY-MM-DD HH24:MI:SS') as update_date
	,(select concat(first_name_th,' ',last_name_th) from m_employee where employee_id = mc.delete_by) as delete_by
	,to_char(mc.delete_date,'YYYY-MM-DD HH24:MI:SS') as delete_date
	,ms.setup_data->>'calculate_enable' as calculate_enable
	,ms.setup_data->>'price_of_cardloss' as price_of_cardloss
	,ms.setup_data->>'except_time_split_from_day' as except_time_split_from_day
	,ms2.setup_data->>'booking_estamp_verify' as booking_estamp_verify
	,ms2.setup_data->>'visitor_estamp_verify' as visitor_estamp_verify
	,CASE WHEN ms3.setup_data->>'line_notification_mode' = 'broadcast' THEN true ELSE false END  as line_notification_broadcast
	,mc.line_company_data
	from m_company mc
	left join m_setup ms
	on mc.company_id = ms.company_id
	left join m_setup ms2
	on ms.company_id = ms2.company_id
	left join m_setup ms3
	on ms.company_id = ms3.company_id
	where mc.company_id = @company_id
	and ms.ref_setup_id = 8
	and ms2.ref_setup_id = 3
	and ms3.ref_setup_id = 1
	 limit 1;`
	rows, err := db.GetDB().Raw(query, sql.Named("company_id", companyRequestDb.Company_id)).Rows()

	if err != nil {
		fmt.Printf("Get by id company error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetErrorLogCompanyFile(), nil, fmt.Sprintf("Get by id company failed : %s", err))
		defer rows.Close()
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&companyResponseDb)
			db.GetDB().ScanRows(rows, &companyResponseDb)
			// do something
		}
		if companyResponseDb.Company_id == 0 {
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCompanyNotInBase})
			utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), nil, "Get by id company Not In Base.")
			return
		}
		resCompanyInfo := convetDataFromDbToResData(companyResponseDb)
		fmt.Printf("Get by id company successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": resCompanyInfo, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), nil, "Get by id company successfully.")
	}
}

func convetDataFromDbToResData(companyResponseDb model_company.CompanyGetByIdFromDBResponse) model_company.CompanyGetByIdResponse {
	var lineCompanyDataMap map[string]interface{}
	json.Unmarshal([]byte(companyResponseDb.Line_company_data), &lineCompanyDataMap)
	companyRes := model_company.CompanyGetByIdResponse{
		Company_id:                  companyResponseDb.Company_id,
		Company_code:                companyResponseDb.Company_code,
		Company_name:                companyResponseDb.Company_name,
		Company_promotion:           companyResponseDb.Company_promotion,
		Company_start_date:          companyResponseDb.Company_start_date,
		Company_expire_date:         companyResponseDb.Company_expire_date,
		Status:                      companyResponseDb.Status,
		Company_remark:              companyResponseDb.Company_remark,
		Create_by:                   companyResponseDb.Create_by,
		Create_date:                 companyResponseDb.Create_date,
		Update_by:                   companyResponseDb.Update_by,
		Update_date:                 companyResponseDb.Update_date,
		Delete_by:                   companyResponseDb.Delete_by,
		Delete_date:                 companyResponseDb.Delete_date,
		Calculate_enable:            companyResponseDb.Calculate_enable,
		Price_of_cardloss:           companyResponseDb.Price_of_cardloss,
		Except_time_split_from_day:  companyResponseDb.Except_time_split_from_day,
		Booking_estamp_verify:       companyResponseDb.Booking_estamp_verify,
		Visitor_estamp_verify:       companyResponseDb.Visitor_estamp_verify,
		Line_notification_broadcast: companyResponseDb.Line_notification_broadcast,
		Line_company_data:           lineCompanyDataMap,
	}
	return companyRes
}
