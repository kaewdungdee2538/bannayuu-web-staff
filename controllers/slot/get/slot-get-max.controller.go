package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_slot_get "bannayuu-web-admin/model/slot/get"
	"bannayuu-web-admin/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSlotMax(c *gin.Context) {
	var request model_slot_get.SlotGetRequest
	var response model_slot_get.SlotGetResponse
	if err := c.ShouldBind(&request); err != nil {
		fmt.Printf("Combine Error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		utils.WriteLogInterface(utils.GetErrorLogGetSlotFile(), nil, fmt.Sprintf("Get slot max failed : %s", err))
		return
	}
	query := `
	SELECT
		visitor_slot_id
		,visitor_slot_number
		,CASE WHEN vs.status_flag = 'Y' THEN true ELSE false END AS use_status	
		,vs.company_id
		,mc.company_code
		,mc.company_name
	FROM m_visitor_slot vs
	LEFT JOIN m_company mc 
	ON vs.company_id = mc.company_id
	WHERE vs.company_id = @company_id
	ORDER BY visitor_slot_number DESC
	LIMIT 1
	;`

	rows, err := db.GetDB().Raw(query,
		sql.Named("company_id", request.Company_id),
	).Rows()

	if err != nil {
		msgErr := fmt.Sprintf("Get slot max error : %s", err)
		fmt.Println(msgErr)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetErrorLogGetSlotFile(), nil, fmt.Sprintf(msgErr))
		defer rows.Close()
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&response)
			db.GetDB().ScanRows(rows, &response)
			// do something
		}

		c.JSON(http.StatusOK, gin.H{"error": false, "result": response, "message": constants.MessageSuccess})
	}
}
