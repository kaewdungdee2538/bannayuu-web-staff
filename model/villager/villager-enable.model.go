package model

import (
	"encoding/json"
)

type VillageEnableRequest struct {
	Home_line_id int    `json:"home_line_id"`
	Remark       string `json:"remark"`
}

func (res VillageEnableRequest) ConvertStructToJson() string {
	text, err := json.Marshal(res)
	if err != nil {
		return ""
	}
	return string(text)
}

type VillageEnableResponse struct {
	Home_line_id int `json:"home_line_id"`
}

func (res VillageEnableResponse) ConvertStructToJson() string {
	text, err := json.Marshal(res)
	if err != nil {
		return ""
	}
	return string(text)
}
