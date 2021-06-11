package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	home_model "bannayuu-web-admin/model/home"
	villager_model "bannayuu-web-admin/model/villager"
	"bannayuu-web-admin/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddVillagerArray(c *gin.Context) {
	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("jwt_employee_id : %v \n", jwtemployeeid)
	var villagerModelArrReq villager_model.VillagerAddRequestModel

	if err := c.ShouldBind(&villagerModelArrReq); err == nil {
		fmt.Printf("villager Model : %v \n", villagerModelArrReq)
		//----------Insert or Update villager Address
		for i, value := range villagerModelArrReq.Data {
			fmt.Println(i, value)
			//---------------Middleware
			if err, messageMiddleware := addVillagerMiddleware(&value); err {
				c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": messageMiddleware})
				utils.WriteLogInterface(utils.GetAccessLogVillagerFile(), nil, messageMiddleware)
				return
			}
			//-----------------Check Home Address
			if isHave, home_id := checkHomeInDataBaseForVillager(villagerModelArrReq.Company_id, &value); isHave {
				//----------------Check Villager
				if isHaveVillager, home_line_id := checkVillagerInDataBaseForVillager(villagerModelArrReq.Company_id, &value); isHaveVillager {
					//----------------Update villager address
					if err, message := updateVillagerQuery(&value, villagerModelArrReq.Company_id, jwtemployeeid, home_id, home_line_id); err {
						c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": message})
						return
					}
				} else {
					//----------------Insert villager address
					if err, message := insertVillagerQuery(&value, villagerModelArrReq.Company_id, jwtemployeeid, home_id); err {
						c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": message})
						return
					}
				}
			} else {
				c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": home_id})
				return
			}
		}
		//----------When Not Error
		fmt.Printf("add villager successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogVillagerFile(), nil, "Add Villager successfully.")
	} else {
		utils.WriteLog(utils.GetAccessLogVillagerFile(), constants.MessageCombineFailed)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	}
}

func checkHomeInDataBaseForVillager(company_id string, villager_obj *villager_model.VillagerRequestModel) (bool, string) {
	var HomeIdResponseDb home_model.HomeIDResponseDb
	query := `select home_id from m_home where company_id = @company_id and home_address = @home_address`
	rows, err := db.GetDB().Raw(query,
		sql.Named("company_id", company_id),
		sql.Named("home_address", villager_obj.Home_address)).Rows()

	if err != nil {
		defer rows.Close()
		message := fmt.Sprintf("Get by id home error : %s\n", err)
		utils.WriteLogInterface(utils.GetAccessLogVillagerFile(), nil, fmt.Sprintf("Get by id home failed : %s", err))
		fmt.Println(message)
		return false, message
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&HomeIdResponseDb)
			db.GetDB().ScanRows(rows, &HomeIdResponseDb)
			// do something
		}
		if HomeIdResponseDb.Home_id == 0 {
			message := fmt.Sprintf("ไม่พบบ้านเลขที่ %s ในระบบ\n", villager_obj.Home_address)
			utils.WriteLogInterface(utils.GetAccessLogVillagerFile(), nil, message)
			fmt.Println(message)
			return false, message
		}
		fmt.Printf("Get by id at home address %s In Base\n", villager_obj.Home_address)
		// utils.WriteLogInterface(utils.GetAccessLogVillagerFile(), nil, "Get by id home successfully.")
		return true, fmt.Sprint(HomeIdResponseDb.Home_id)
	}
}

func checkVillagerInDataBaseForVillager(company_id string, villager_obj *villager_model.VillagerRequestModel) (bool, int) {
	var HomeLineIdResponseDb villager_model.VillagerIDResponseDb
	query := `select home_line_id from m_home_line where company_id = @company_id and home_line_mobile_phone = @tel_number`
	rows, err := db.GetDB().Raw(query,
		sql.Named("company_id", company_id),
		sql.Named("tel_number", villager_obj.Tel_number),
	).Rows()

	if err != nil {
		defer rows.Close()
		message := fmt.Sprintf("Get by id home line error : %s\n", err)
		utils.WriteLogInterface(utils.GetAccessLogVillagerFile(), nil, fmt.Sprintf("Get by id home line failed : %s", err))
		fmt.Println(message)
		return false, 0
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&HomeLineIdResponseDb)
			db.GetDB().ScanRows(rows, &HomeLineIdResponseDb)
			// do something
		}
		if HomeLineIdResponseDb.Home_line_id == 0 {
			message := fmt.Sprintf("Get by id at home line %s Not In Base.\n", villager_obj.Home_address)
			utils.WriteLogInterface(utils.GetAccessLogVillagerFile(), nil, message)
			fmt.Println(message)
			return false, 0
		}
		fmt.Printf("Get by id at home line %s In Base\n", villager_obj.Home_address)
		// utils.WriteLogInterface(utils.GetAccessLogVillagerFile(), nil, "Get by id home successfully.")
		return true, HomeLineIdResponseDb.Home_line_id
	}
}

func mapLogVillagerDataToJsonString(
	villager_obj *villager_model.VillagerRequestModel,
	company_id string,
	jwtemployeeid interface{},
	home_id string) (bool, string) {
	req_map := map[string]interface{}{
		"home_id":      home_id,
		"home_address": villager_obj.Home_address,
		"first_name":   villager_obj.First_name,
		"last_name":    villager_obj.Last_name,
		"tel_number":   villager_obj.Tel_number,
		"remark":       villager_obj.Remark,
		"company_id":   company_id,
		"create_by":    jwtemployeeid,
	}
	err, setup_data := utils.ConvertInterfaceToJSON(req_map)
	if err {
		return true, ""
	}
	return false, setup_data
}

func mapLogVillagerUpdateDataToJsonString(
	villager_obj *villager_model.VillagerRequestModel,
	company_id string,
	jwtemployeeid interface{},
	home_id string,
	home_line_id int) (bool, string) {
	req_map := map[string]interface{}{
		"home_line_id": home_line_id,
		"home_id":      home_id,
		"home_address": villager_obj.Home_address,
		"first_name":   villager_obj.First_name,
		"last_name":    villager_obj.Last_name,
		"tel_number":   villager_obj.Tel_number,
		"remark":       villager_obj.Remark,
		"company_id":   company_id,
		"create_by":    jwtemployeeid,
	}
	err, setup_data := utils.ConvertInterfaceToJSON(req_map)
	if err {
		return true, ""
	}
	return false, setup_data
}
