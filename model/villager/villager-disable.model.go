package model

import (
	"encoding/json"
)

type VillageDisableRequest struct {
	Home_line_id int    `json:"home_line_id"`
	Remark       string `json:"remark"`
}

func (res VillageDisableRequest) ConvertStructToJson() string {
	text, err := json.Marshal(res)
	if err != nil {
		return ""
	}
	return string(text)
}

type VillageDisableResponse struct {
	Home_line_id int `json:"home_line_id"`
}

func (res VillageDisableResponse) ConvertStructToJson() string {
	text, err := json.Marshal(res)
	if err != nil {
		return ""
	}
	return string(text)
}
