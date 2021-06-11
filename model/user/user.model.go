package model

type UserAddRequestModel struct {
	First_name            string
	Last_name             string
	Address               string
	Mobile                string
	Line                  string
	Email                 string
	Username              string
	Password              string
	Employee_privilege_id string
	Status                string
	Employee_type         string
	Company_id            string
	Company_list          []int
}

type UserGetRequestModel struct {
	Company_id string
	Full_name  string
}

type UserGetResponseModel struct {
	Employee_id             int    `json:"employee_id"`
	Employee_code           string `json:"employee_code"`
	First_name_th           string `json:"first_name_th"`
	Last_name_th            string `json:"last_name_th"`
	Username                string `json:"username"`
	Employee_privilege_type string `json:"employee_privilege_type"`
}

type UserGetByIdRequestModel struct {
	Company_id  string
	Employee_id string
}

type UserInfoGetResponseModel struct {
	Employee_id                string `json:"employee_id"`
	Employee_code              string `json:"employee_code"`
	First_name_th              string `json:"first_name_th"`
	Last_name_th               string `json:"last_name_th"`
	Address                    string `json:"address"`
	Employee_telephone         string `json:"employee_telephone"`
	Employee_mobile            string `json:"employee_mobile"`
	Employee_line              string `json:"employee_line"`
	Username                   string `json:"username"`
	Remark                     string `json:"remark"`
	Employee_privilege_name_th string `json:"employee_privilege_name_th"`
	Employee_privilege_type    string `json:"employee_privilege_type"`
	Create_by                  string `json:"create_by"`
	Create_date                string `json:"create_date"`
	Update_by                  string `json:"update_by"`
	Update_date                string `json:"update_date"`
	Delete_by                  string `json:"delete_by"`
	Delete_date                string `json:"delete_date"`
	Delete_flag                string `json:"delete_flag"`
}
