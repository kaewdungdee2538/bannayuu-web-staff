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
}
