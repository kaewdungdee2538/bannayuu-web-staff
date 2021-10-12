package model

type UserResetPasswordRequestModel struct {
	Employee_id string
	Hold_time   string
	Remark      string
	Company_id  string
}

type UserResetPasswordResponseModel struct {
	Tepi_id     int    `json:"tepi_id"`
	Tepi_code   string `json:"tepi_code"`
	Employee_id int    `json:"employee_id"`
	Company_id  int    `json:"company_id"`
}
