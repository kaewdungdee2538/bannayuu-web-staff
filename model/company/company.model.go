package model

import "mime/multipart"

type CompanyAddModelRequest struct {
	Image                      *multipart.FileHeader `form:"image" binding:"required"`
	Company_code               string                `form:"company_code" binding:"required"`
	Company_name               string                `form:"company_name" binding:"required"`
	Company_promotion          string                `form:"company_promotion" binding:"required"`
	Company_start_date         string                `form:"company_start_date" binding:"required"`
	Company_expire_date        string                `form:"company_expire_date" binding:"required"`
	Calculate_enable           bool                  `form:"calculate_enable"`
	Except_time_split_from_day bool                  `form:"except_time_split_from_day"`
	Price_of_cardloss          int                   `form:"price_of_cardloss"`
	Booking_verify             string                `form:"booking_verify"`
	Visitor_verify             string                `form:"visitor_verify"`
	Booking_estamp_verify      bool                  `form:"booking_estamp_verify"`
	Visitor_estamp_verify      bool                  `form:"visitor_estamp_verify"`
}

type CompanyEditModelRequest struct {
	Image                      *multipart.FileHeader `form:"image" binding:"required"`
	Company_id                 string                `form:"company_id" binding:"required"`
	Company_code               string                `form:"company_code" binding:"required"`
	Company_name               string                `form:"company_name" binding:"required"`
	Company_promotion          string                `form:"company_promotion" binding:"required"`
	Company_start_date         string                `form:"company_start_date" binding:"required"`
	Company_expire_date        string                `form:"company_expire_date" binding:"required"`
	Remark                     string                `form:"remark"`
	Calculate_enable           bool                  `form:"calculate_enable"`
	Except_time_split_from_day bool                  `form:"except_time_split_from_day"`
	Price_of_cardloss          int                   `form:"price_of_cardloss"`
	Booking_verify             string                `form:"booking_verify"`
	Visitor_verify             string                `form:"visitor_verify"`
	Booking_estamp_verify      bool                  `form:"booking_estamp_verify"`
	Visitor_estamp_verify      bool                  `form:"visitor_estamp_verify"`
}

type CompanyDisableModelRequest struct {
	Image      *multipart.FileHeader `form:"image" binding:"required"`
	Company_id string                `form:"company_id" binding:"required"`
	Remark     string                `form:"remark"`
}

type CompanyGetAllRequest struct {
	Company_code_or_name string
}
type CompanyGetAllResponse struct {
	Company_id        int    `json:"company_id"`
	Company_code      string `json:"company_code"`
	Company_name      string `json:"company_name"`
	Company_promotion string `json:"company_promotion"`
	Status            string `json:"status"`
}

type CompanyGetByIdResquest struct {
	Company_id int
}

type CompanyGetByIdResponse struct {
	Company_id                 int    `json:"company_id"`
	Company_code               string `json:"company_code"`
	Company_name               string `json:"company_name"`
	Company_promotion          string `json:"company_promotion"`
	Status                     string `json:"status"`
	Company_start_date         string `json:"company_start_date"`
	Company_expire_date        string `json:"company_expire_date"`
	Company_remark             string `json:"company_remark"`
	Create_by                  string `json:"create_by"`
	Create_date                string `json:"create_date"`
	Update_by                  string `json:"update_by"`
	Update_date                string `json:"update_date"`
	Delete_by                  string `json:"delete_by"`
	Delete_date                string `json:"delete_date"`
	Calculate_enable           bool   `json:"calculate_enable"`
	Price_of_cardloss          int    `json:"price_of_cardloss"`
	Except_time_split_from_day bool   `json:"except_time_split_from_day"`
	Booking_estamp_verify      bool   `json:"booking_estamp_verify"`
	Visitor_estamp_verify      bool   `json:"visitor_estamp_verify"`
}


type CompanyGetIdNotDisableResponseModel struct {
	Company_id int `json:"company_id"`
}
