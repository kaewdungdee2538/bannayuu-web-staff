package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	home_model "bannayuu-web-admin/model/home"
	"bannayuu-web-admin/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func AddHomeArray(c *gin.Context) {
	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("jwt_employee_id : %v \n", jwtemployeeid)
	var homeModelArrReq home_model.HomeAddRequestModel

	if err := c.ShouldBind(&homeModelArrReq); err == nil {
		fmt.Printf("home Model : %v \n", homeModelArrReq)
		//----------Insert or Update Home Address
		for i, value := range homeModelArrReq.Data {
			fmt.Println(i, value)
			if err, messageMiddleware := insertHomeMiddleware(&value); err {
				c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": messageMiddleware})
				utils.WriteLogInterface(utils.GetAccessLogHomeFile(), nil, messageMiddleware)
				return
			}
			if isHave := checkHomeInDataBase(homeModelArrReq.Company_id, &value); !isHave {
				//----------------Insert home address
				if err, message := insertHomeQuery(&value, homeModelArrReq.Company_id, jwtemployeeid); err {
					utils.WriteLogInterface(utils.GetAccessLogHomeFile(), nil, message)
					c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": message})
					return
				}
			}
		}
		//----------When Not Error
		fmt.Printf("add Home successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogHomeFile(), nil, "Add Home successfully.")
	} else {
		utils.WriteLog(utils.GetAccessLogHomeFile(), constants.MessageCombineFailed)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	}
}

func checkHomeInDataBase(company_id string, home_obj *home_model.HomeRequestModel) bool {
	var HomeIdResponseDb home_model.HomeIDResponseDb
	query := `select home_id from m_home where company_id = @company_id and home_address = @home_address`
	rows, err := db.GetDB().Raw(query,
		sql.Named("company_id", company_id),
		sql.Named("home_address", home_obj.Home_address)).Rows()

	if err != nil {
		fmt.Printf("Get by id home error : %s\n", err)
		utils.WriteLogInterface(utils.GetAccessLogHomeFile(), nil, fmt.Sprintf("Get by id home failed : %s", err))
		defer rows.Close()
		return false
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&HomeIdResponseDb)
			db.GetDB().ScanRows(rows, &HomeIdResponseDb)
			// do something
		}
		if HomeIdResponseDb.Home_id == 0 {
			message := fmt.Sprintf("Get by id at home address %s Not In Base.\n", home_obj.Home_address)
			utils.WriteLogInterface(utils.GetAccessLogHomeFile(), nil, message)
			fmt.Println(message)
			return false
		}
		fmt.Printf("Get by id at home address %s In Base\n", home_obj.Home_address)
		// utils.WriteLogInterface(utils.GetAccessLogHomeFile(), nil, "Get by id home successfully.")
		return true
	}
}

func insertHomeQuery(
	homeModelReq *home_model.HomeRequestModel,
	company_id_str string,
	jwtemployeeid interface{},
) (bool, string) {
	fmt.Println("Insert Home Address")
	// home_code := fmt.Sprintf("fun_generate_uuid('HOME'||trim(to_char(%s,'000')),6)", company_id)
	company_id, _ := strconv.ParseInt(company_id_str, 10, 64)
	err_log_obj, log_obj := mapLogHomeDataToJsonString(homeModelReq, company_id_str, jwtemployeeid)
	if err_log_obj {
		utils.WriteLog(utils.GetAccessLogHomeFile(), constants.MessageCovertObjTOJSONFailed)
	}
	//---------Calcualte setup object
	query := fmt.Sprintf(`
		with hometb as (
		insert into m_home(
			home_code
			,home_address
			,home_remark
			,create_by
			,create_date
			,company_id
		) values(
			fun_generate_uuid('HOME'||trim(to_char(%v,'000')),6)
			,@home_address
			,@remark
			,@create_by
			,current_timestamp
			,@company_id
			)
			RETURNING home_id as homeid
		)
		insert into log_home(
				lh_code
				,lh_name
				,lh_data
				,lh_type
				,create_by
				,create_date
				,company_id
			)values(
				fun_generate_uuid('LH'||trim(to_char(%v,'000'))||(select homeid::text FROM hometb),5)
				,'เพิ่มบ้านหลังใหม่'
				,@log_data
				,'IMPORT'
				,@create_by
				,current_timestamp
				,@company_id
			)
			`, company_id, company_id)

	values := map[string]interface{}{
		"home_address": homeModelReq.Home_address,
		"remark":       homeModelReq.Remark,
		"create_by":    fmt.Sprint(jwtemployeeid),
		"company_id":   company_id,
		"log_data":     log_obj}

	if err, message := db.SaveTransactionDB(query, values); err {
		fmt.Printf("add home Failed")
		utils.WriteLogInterface(utils.GetAccessLogHomeFile(), values, fmt.Sprintf("Add Home Failed : %s", message))
		return true, constants.MessageFailed
	} else {
		fmt.Printf("add home successfully")
		utils.WriteLogInterface(utils.GetAccessLogHomeFile(), values, "Add Home successfully.")
		return false, constants.MessageSuccess
	}
}

func mapLogHomeDataToJsonString(homeModelReq *home_model.HomeRequestModel,
	company_id string,
	jwtemployeeid interface{}) (bool, string) {
	req_map := map[string]interface{}{
		"home_address": homeModelReq.Home_address,
		"remark":       homeModelReq.Remark,
		"company_id":   company_id,
		"create_by":    jwtemployeeid,
	}
	err, setup_data := utils.ConvertInterfaceToJSON(req_map)
	if err {
		return true, ""
	}
	return false, setup_data
}

func insertHomeMiddleware(homeModelReq *home_model.HomeRequestModel) (bool, string) {
	Home_address := strings.TrimSpace(homeModelReq.Home_address)
	if len(Home_address) == 0 {
		return true, constants.MessageHomeAddressNotFound
	} else if utils.IsNotStringAlphabetRemark(Home_address) {
		return true, fmt.Sprintf("%s ,[ที่อยู่ : %s]", constants.MessageHomeAddressProhibitSpecial, Home_address)
	} else if utils.IsNotStringAlphabetRemark(homeModelReq.Remark) {
		return true, constants.MessageRemarkProhibitSpecial
	}
	return false, ""
}
