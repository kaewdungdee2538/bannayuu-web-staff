package models

import (
	"encoding/json"
)

type SlotGetRequest struct {
	Company_id int `json:"company_id"`
}

func (res SlotGetRequest) ConvertStructToJson() string {
	text, err := json.Marshal(res)
	if err != nil {
		return ""
	}
	return string(text)
}

type SlotGetResponse struct {
	Visitor_slot_id     int    `json:"visitor_slot_id"`
	Visitor_slot_number int    `json:"visitor_slot_number"`
	Visitor_slot_code   string `json:"visitor_slot_code"`
	Use_status          bool   `json:"use_status"`
	Company_id          int    `json:"company_id"`
	Company_name        string `json:"company_name"`
}

func (res SlotGetResponse) ConvertStructToJson() string {
	text, err := json.Marshal(res)
	if err != nil {
		return ""
	}
	return string(text)
}

type SlotGetMaxResponse struct {
	Visitor_slot_id     int    `json:"visitor_slot_id"`
	Visitor_slot_number int    `json:"visitor_slot_number"`
	Visitor_slot_code   string `json:"visitor_slot_code"`
	Use_status          bool   `json:"use_status"`
	Company_id          int    `json:"company_id"`
	Company_code        string `json:"company_code"`
	Company_name        string `json:"company_name"`
}

func (res SlotGetMaxResponse) ConvertStructToJson() string {
	text, err := json.Marshal(res)
	if err != nil {
		return ""
	}
	return string(text)
}
