package interceptor

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	villager_model "bannayuu-web-admin/model/villager"
	format_utls "bannayuu-web-admin/utils"
	"database/sql"
	"errors"
)
func checkValueHomeLineId(homeLineId string) error {
	if len(homeLineId) == 0 {
		return errors.New(constants.MessageHomeLineIdNotFound) 
	} else if format_utls.IsNotStringNumber(homeLineId) {
		return  errors.New(constants.MessageHomeLineIdNotNumber)
	}
	return checkValuesGetHomeLineId(homeLineId)
}

func checkValuesGetHomeLineId(homeLineId string) error {
	var res villager_model.VillageDisableResponse

	query := `select home_line_id from m_home_line
	where home_line_id = @home_line_id;
	`
	rows, _ := db.GetDB().Raw(query, sql.Named("home_line_id", homeLineId)).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&res)
		db.GetDB().ScanRows(rows, &res)
		// do something
	}
	if res.Home_line_id == 0 {
		return errors.New(constants.MessageHomeLineNotInBase)
	}
	return nil
}
