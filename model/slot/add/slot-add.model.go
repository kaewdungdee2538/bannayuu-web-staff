package models

import (
	"encoding/json"
)

type SlotAddRequest struct {
	Slot_count      int    `json:"slot_count"`
	Company_id      int    `json:"company_id"`
	Guardhouse_id   int    `json:"guardhouse_id"`
	Guardhouse_code string `json:"guardhouse_code"`
}

func (res SlotAddRequest) ConvertStructToJson() string {
	text, err := json.Marshal(res)
	if err != nil {
		return ""
	}
	return string(text)
}

type SlotAddResponse struct {
	Func_addvisitorslot bool `json:"func_addvisitorslot"`
}

func (res SlotAddResponse) ConvertStructToJson() string {
	text, err := json.Marshal(res)
	if err != nil {
		return ""
	}
	return string(text)
}
