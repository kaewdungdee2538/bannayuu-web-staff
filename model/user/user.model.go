package model

import "mime/multipart"

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

type UserEditInfoRequestModel struct {
	Employee_id string
	First_name  string
	Last_name   string
	Address     string
	Mobile      string
	Line        string
	Email       string
	Company_id  string
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

type UserChangePrivilegeRequestModel struct {
	Image                 *multipart.FileHeader `form:"image" binding:"required"`
	Company_id            string                `form:"company_id"`
	Employee_id           string                `form:"employee_id"`
	Remark                string                `form:"remark" `
	Employee_privilege_id string                `form:"employee_privilege_id"`
	Employee_type         string                `form:"employee_type"`
}

type UserChangeMainCompanyRequestModel struct {
	Image          *multipart.FileHeader `form:"image" binding:"required"`
	Employee_id    string                `form:"employee_id"`
	Old_company_id string                `form:"old_company_id"`
	New_company_id string                `form:"new_company_id"`
	Remark         string                `form:"remark" `
}

type UserAddOrDeleteCompanyListRequestModel struct {
	Image        *multipart.FileHeader `form:"image" binding:"required"`
	Employee_id  string                `form:"employee_id"`
	Company_id   string                `form:"company_id"`
	Company_list []string              `form:"company_list"`
	Remark       string                `form:"remark" `
}

type UserGetPrivilegeResponseModel struct {
	Employee_privilege_id      int    `json:"employee_privilege_id"`
	Employee_privilege_name_th string `json:"employee_privilege_name_th"`
	Employee_privilege_type    string `json:"employee_privilege_type"`
}
