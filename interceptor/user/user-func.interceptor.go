package interceptor

import (
	db "bannayuu-web-admin/db"
	constants "bannayuu-web-admin/constants"
	"database/sql"
	"strings"
)

type UserInterceptorModel struct {
	Employee_id   int
}

func checkUserIsDuplicateInbaseWhenCreateUser(typeUser string, userName string, comapnyId string) (bool, string) {
	switch strings.ToUpper(typeUser) {
	case "USER":
		return checkUserIsDuplicateInBaseNormalTypeByCompanyId(userName, comapnyId)
	default:
		return checkUserIsDuplicateInBaseManagementType(userName)
	}
}

func checkUserIsDuplicateInBaseNormalTypeByCompanyId(userName string, comapnyId string) (bool, string) {
	var userModel UserInterceptorModel
	query := `select employee_id from m_employee me
	left join m_employee_privilege mep
	on me.employee_privilege_id = mep.employee_privilege_id
	where username = @username and me.company_id = @company_id
	and mep.employee_privilege_type = 'USER'
	`;
	rows, _ := db.GetDB().Raw(query,sql.Named("username",userName),sql.Named("company_id",comapnyId)).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&userModel)
		db.GetDB().ScanRows(rows, &userModel)
		// do something
	}
	if userModel.Employee_id > 0 {
		return true, constants.MessageUserIsDuplicate
	}
	return false, ""
}

func checkUserIsDuplicateInBaseManagementType(userName string) (bool, string) {
	var userModel UserInterceptorModel
	query := `select employee_id from m_employee me
	left join m_employee_privilege mep
	on me.employee_privilege_id = mep.employee_privilege_id
	where username = @username
	and mep.employee_privilege_type != 'USER'
	`;
	rows, _ := db.GetDB().Raw(query,sql.Named("username",userName)).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&userModel)
		db.GetDB().ScanRows(rows, &userModel)
		// do something
	}
	if userModel.Employee_id > 0 {
		return true, constants.MessageUserIsDuplicate
	}
	return false, ""
}


func checkEmployeeIdInBase(employeeId string,companyId string) (bool,string){
	var userModel UserInterceptorModel
	query := `select employee_id from m_employee me
	left join m_employee_privilege mep
	on me.employee_privilege_id = mep.employee_privilege_id
	where employee_id = @employee_id and me.company_id = @company_id
	`;
	rows, _ := db.GetDB().Raw(query,sql.Named("employee_id",employeeId),sql.Named("company_id",companyId)).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&userModel)
		db.GetDB().ScanRows(rows, &userModel)
		// do something
	}
	if userModel.Employee_id == 0 {
		return true, constants.MessageUserNotInBase
	}
	return false, ""
}