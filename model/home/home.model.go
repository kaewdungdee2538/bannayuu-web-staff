package model

type HomeAddRequestModel struct {
	Company_id string
	Data       []HomeRequestModel
}

type HomeRequestModel struct {
	Home_address string
	Remark       string
}

type HomeIDResponseDb struct {
	Home_id int
}

type HomeGetRequestModel struct {
	Company_id   string
	Home_address string
}

type HomeGetResponseModel struct {
	Home_id      int    `json:"home_id"`
	Home_code    string `json:"home_code"`
	Home_address string `json:"home_address"`
	Home_remark  string `json:"home_remark"`
	Status       string `json:"status"`
	Create_by    string `json:"create_by"`
	Create_date  string `json:"create_date"`
}
