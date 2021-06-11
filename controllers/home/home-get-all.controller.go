package controller

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_home "bannayuu-web-admin/model/home"
	"bannayuu-web-admin/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHomeAll(c *gin.Context) {
	var homeRequest model_home.HomeGetRequestModel
	var homeResponse []model_home.HomeGetResponseModel
	
	if err := c.ShouldBind(&homeRequest); err != nil {
		fmt.Printf("Combine Error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
		return
	}
	query := `select home_id,home_code,home_address,home_remark,
	case when mh.delete_flag = 'Y' then 'DISABLE' else
	'NORMAL' end as status,
	concat(me.first_name_th,' ',me.last_name_th) as create_by,
	to_char(mh.create_date,'YYYY-MM-DD HH24:MI:SS') as create_date
	from m_home mh
	left join m_employee me
	on mh.create_by = me.employee_id::varchar
	where mh.company_id = @company_id
	and home_address LIKE @home_address
	order by home_address
	;`

	likeStr := "%"
	home_address := fmt.Sprintf("%s%s%s", likeStr, homeRequest.Home_address, likeStr)

	rows, err := db.GetDB().Raw(query,
		sql.Named("company_id", homeRequest.Company_id),
		sql.Named("home_address", home_address),
		).Rows()

	if err != nil {
		fmt.Printf("Get Home All By Company ID error : %s", err)
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageFailed})
		utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), nil, fmt.Sprintf("Get Home All By Company ID failed : %s", err))
		defer rows.Close()
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&homeResponse)
			db.GetDB().ScanRows(rows, &homeResponse)
			// do something
		}

		fmt.Printf("Get Home All By Company ID successfully")
		c.JSON(http.StatusOK, gin.H{"error": false, "result": homeResponse, "message": constants.MessageSuccess})
		utils.WriteLogInterface(utils.GetAccessLogCompanyFile(), nil, "Get Home All By Company ID successfully.")
	}
}
