package interceptor

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	"database/sql"
	"strings"
)

type UserInterceptorModel struct {
	Employee_id   int
}

type UserPrivilegeInterceptorModel struct {
	Employee_privilege_id   int
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

func checkEmployeePrivilegeIdForCustomerInBase(employeePrivilegeId string,employeeType string) (bool,string){
	var userPrivilegeModel UserPrivilegeInterceptorModel
	query := `select employee_privilege_id from m_employee_privilege 
	where employee_privilege_id = @employee_privilege_id and employee_privilege_type = @employee_type
	and employee_privilege_id not in (6,7)`;
	rows, _ := db.GetDB().Raw(query,sql.Named("employee_privilege_id",employeePrivilegeId),sql.Named("employee_type",employeeType)).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&userPrivilegeModel)
		db.GetDB().ScanRows(rows, &userPrivilegeModel)
		// do something
	}
	if userPrivilegeModel.Employee_privilege_id == 0 {
		return true, constants.MessageUserPrivilegeNotInBase
	}
	return false, ""
}