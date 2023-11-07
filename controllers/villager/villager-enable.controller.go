package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_villager "bannayuu-web-admin/model/villager"
	"bannayuu-web-admin/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VillagerEnableById(c *gin.Context) {
	jwtemployeeid, _ := c.Get("jwt_employee_id")

	var request model_villager.VillageEnableRequest
	var response model_villager.VillageEnableResponse
	if err := c.ShouldBind(&request); err != nil {
		fmt.Printf("Combine Error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		return
	}
	query := `UPDATE m_home_line SET
		delete_flag = 'N'
		,update_date = current_timestamp
		,update_by = @update_by
		,home_line_remark = @remark
	WHERE home_line_id = @home_line_id
	RETURNING home_line_id
	;`
	rows, err := db.GetDB().Raw(query,
		sql.Named("home_line_id", request.Home_line_id),
		sql.Named("remark", request.Remark),
		sql.Named("update_by", fmt.Sprint(jwtemployeeid)),
	).Rows()

	if err != nil {
		msgErr := fmt.Sprintf("Enable Villager error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		fmt.Println(msgErr)
		utils.WriteLogInterface(utils.GetErrorLogVillagerEnableFile(), nil, msgErr)
		defer rows.Close()
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&response)
			db.GetDB().ScanRows(rows, &response)
			// do something
		}
		if response.Home_line_id == 0 {
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCompanyNotInBase})
			utils.WriteLogInterface(utils.GetAccessLogVillagerEnableFile(), nil, "Enable Villager Failed.")
			return
		}

		c.JSON(http.StatusOK, gin.H{"error": false, "result": response, "message": constants.MessageSuccess})
	}
}
