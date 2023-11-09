package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	slot_model "bannayuu-web-admin/model/slot/add"
	"bannayuu-web-admin/utils"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddSlotManual(c *gin.Context) {
	jwtemployeeid, _ := c.Get("jwt_employee_id")
	fmt.Printf("jwt_employee_id : %v \n", jwtemployeeid)
	var slotAddReq slot_model.SlotAddRequest

	if err := c.ShouldBind(&slotAddReq); err == nil {
		fmt.Printf("slot add manual Model : %v \n", slotAddReq)
		//----------Insert User
		if err, message := addSlotManualToDb(slotAddReq, jwtemployeeid); err {
			utils.WriteLogInterface(utils.GetErrorLogAddSlotManualFile(), nil, message)
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": message})
			return
		}
		//----------When Not Error
		fmt.Printf("add slot number successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageSuccess})
	} else {
		utils.WriteLog(utils.GetAccessLogAddSlotManualFile(), constants.MessageCombineFailed)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	}
}

func addSlotManualToDb(slotAddReq slot_model.SlotAddRequest, jwtemployeeid interface{}) (bool, string) {
	
	var res slot_model.SlotAddResponse

	query := `
	WITH input_data AS (
		SELECT
			$1::BIGINT AS slot_count,
			$2::TEXT AS slot_prefix,
			$3::INTEGER AS company_id,
			$4::INTEGER AS guardhouse_id,
			$5::TEXT AS guardhouse_code
	)
	,select_max_slot_number AS (
		SELECT
			MAX(visitor_slot_number)+1 AS max_slot_number
			,mc.company_id
			,mc.company_code
		FROM m_visitor_slot vs
		LEFT JOIN m_company mc
		ON vs.company_id = mc.company_id
		WHERE vs.company_id = (SELECT company_id FROM input_data)
		GROUP BY mc.company_id
	)
	,select_result AS (
		SELECT
			max_slot_number
			,slot_count
			,(max_slot_number+slot_count)-1 AS slot_number_end
			,slot_prefix
			,ind.company_id
			,company_code
			,guardhouse_id
			,guardhouse_code
		FROM input_data ind
		LEFT JOIN select_max_slot_number msn
		ON ind.company_id = msn.company_id
	)
	SELECT func_addvisitorslot_manual(
		(SELECT slot_number_end FROM select_result)
		,(SELECT slot_prefix FROM select_result)
		,(SELECT company_id FROM select_result)
		,(SELECT company_code FROM select_result)
		,(SELECT guardhouse_id FROM select_result)
		,(SELECT guardhouse_code FROM select_result)
		,(SELECT max_slot_number FROM select_result)
	);
	`
	rows, err := db.GetDB().Raw(query,
		slotAddReq.Slot_count,
		fmt.Sprint(slotAddReq.Company_id),
		slotAddReq.Company_id,
		slotAddReq.Guardhouse_id,
		slotAddReq.Guardhouse_code,
		).Rows()

	if err != nil {
		msgErr := fmt.Sprintf("Add slot number error : %s", err)
		fmt.Println(msgErr)
		utils.WriteLogInterface(utils.GetAccessLogAddSlotManualFile(), nil, msgErr)
		defer rows.Close()
		return true, msgErr
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&res)
			db.GetDB().ScanRows(rows, &res)
			// do something
		}
		if res.Func_addvisitorslot_manual{
			return false, ""
		}
		return true, "เพิ่มเลข Slot ล้มเหลว"
	}
}