package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	villager_model "bannayuu-web-admin/model/villager"
	"bannayuu-web-admin/utils"
	"fmt"
	"strconv"
)

func updateVillagerQuery(
	villager_obj *villager_model.VillagerRequestModel,
	company_id_str string,
	jwtemployeeid interface{},
	home_id string,
	home_line_id int,
) (bool, string) {
	fmt.Println("update Villager Address")
	// home_code := fmt.Sprintf("fun_generate_uuid('HOME'||trim(to_char(%s,'000')),6)", company_id)
	company_id, _ := strconv.ParseInt(company_id_str, 10, 64)
	err_log_obj, log_obj := mapLogVillagerUpdateDataToJsonString(villager_obj, company_id_str, jwtemployeeid, home_id, home_line_id)
	if err_log_obj {
		utils.WriteLog(utils.GetAccessLogVillagerFile(), constants.MessageCovertObjTOJSONFailed)
	}
	//---------Calcualte setup object
	query := fmt.Sprintf(`
		with villagertb as (
		update m_home_line set
			home_line_first_name = @home_line_first_name
			,home_line_last_name = @home_line_last_name
			,home_line_mobile_phone = @home_line_mobile_phone
			,home_line_remark = @remark
			,update_by = @update_by
			,update_date = current_timestamp
			where company_id = @company_id
			and home_line_id = @home_line_id
			RETURNING home_line_id as homelineid
		)
		insert into log_home_line(
				lhl_code
				,lhl_name
				,lhl_data
				,lhl_type
				,create_by
				,create_date
				,company_id
			)values(
				fun_generate_uuid('LHL'||trim(to_char(%v,'000'))||trim(to_char((select homelineid FROM villagertb),'000')),5)
				,'แก้ไขข้อมูลลูกบ้าน'
				,@log_data
				,'UPDATE'
				,@update_by
				,current_timestamp
				,@company_id
			)
			`, company_id)

	values := map[string]interface{}{
		"home_line_id":           home_line_id,
		"home_id":                home_id,
		"home_line_first_name":   villager_obj.First_name,
		"home_line_last_name":    villager_obj.Last_name,
		"home_line_mobile_phone": villager_obj.Tel_number,
		"remark":                 villager_obj.Remark,
		"update_by":              fmt.Sprint(jwtemployeeid),
		"company_id":             company_id,
		"log_data":               log_obj}

	if err, message := db.SaveTransactionDB(query, values); err {
		fmt.Printf("update villager Failed")
		utils.WriteLogInterface(utils.GetAccessLogVillagerFile(), values, fmt.Sprintf("Update Villager Failed : %s", message))
		return true, constants.MessageFailed
	} else {
		fmt.Printf("update villager successfully")
		utils.WriteLogInterface(utils.GetAccessLogVillagerFile(), values, "Update Villager successfully.")
		return false, constants.MessageSuccess
	}
}
