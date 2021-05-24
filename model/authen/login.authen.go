package model

type AuthenUserLoginStc struct {
	Username string
	Password string
}

type AuthenUserLoginForDBStc struct {
	Employee_id     int
	Employee_code   string
	First_name_th   string
	Last_name_th    string
	Username        string
	Password_status bool
	Company_id      string
	Company_name    string
	Company_list    string `gorm:"type:text"`
	Privilege_info  string `gorm:"type:text"`
}

type AuthenUserLoginResponseStc struct {
	Access_token 	string	`json:"access_token"`
	Employee_id     int 	`json:"employee_id"`
	Employee_code   string	`json:"eployee_code"`
	First_name_th   string	`json:"first_name_th"`
	Last_name_th    string	`json:"last_name_th"`
	Username        string	`json:"username"`
	Password_status bool	`json:"password_status"`
	Company_id      string	`json:"company_id"`
	Company_name    string	`json:"company_name"`
	Company_list    interface{}	`json:"company_list"`
	Privilege_info  interface{}	`json:"privilege_info"`
}

