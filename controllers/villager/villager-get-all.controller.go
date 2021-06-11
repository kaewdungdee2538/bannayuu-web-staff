package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_villager "bannayuu-web-admin/model/villager"
	"bannayuu-web-admin/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVillagerAll(c *gin.Context) {
	var villagerRequest model_villager.VillagerGetRequestModel
	var villagerResponse []model_villager.VillagerGetAllResponseModel
	
	if err := c.ShouldBind(&villagerRequest); err != nil {
		fmt.Printf("Combine Error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		return
	}
	query := `select home_line_id as villager_id,home_line_code as villager_code,
	mh.home_address,home_line_first_name as first_name,home_line_last_name as last_name,
	home_line_mobile_phone as tel_number,home_line_remark as remark,
	case when mhl.delete_flag = 'Y' then 'DISABLE' else 'NORMAL' end as status,
	concat(me.first_name_th,' ',me.last_name_th) as create_by,
	to_char(mhl.create_date,'YYYY-MM-DD HH24:MI:SS') as create_date
	from m_home_line mhl
	left join m_home mh
	on mhl.home_id = mh.home_id
	left join m_employee me
	on mhl.create_by = me.employee_id::varchar
	where mhl.company_id = @company_id
	and mh.home_address LIKE @home_address
	and (mhl.home_line_first_name LIKE @full_name or mhl.home_line_last_name LIKE @full_name)
	order by home_address,home_line_first_name,home_line_last_name
	;`

	likeStr := "%"
	home_address := fmt.Sprintf("%s%s%s", likeStr, villagerRequest.Home_address, likeStr)
	full_name := fmt.Sprintf("%s%s%s", likeStr, villagerRequest.Full_name, likeStr)
	rows, err := db.GetDB().Raw(query,
		sql.Named("company_id", villagerRequest.Company_id),
		sql.Named("home_address", home_address),
		sql.Named("full_name", full_name),
		).Rows()

	if err != nil {
		fmt.Printf("Get Villager All By Company ID error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetAccessLogVillagerFile(), nil, fmt.Sprintf("Get Villager All By Company ID failed : %s", err))
		defer rows.Close()
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&villagerResponse)
			db.GetDB().ScanRows(rows, &villagerResponse)
			// do something
		}

		fmt.Printf("Get Villager All By Company ID successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": villagerResponse, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogVillagerFile(), nil, "Get Villager All By Company ID successfully.")
	}
}
