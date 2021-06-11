package model

type VillagerAddRequestModel struct {
	Company_id string
	Data       []VillagerRequestModel
}

type VillagerRequestModel struct {
	Home_address string
	First_name   string
	Last_name    string
	Tel_number   string
	Remark       string
}

type VillagerIDResponseDb struct {
	Home_line_id int
}

type VillagerGetRequestModel struct {
	Company_id   string
	Home_address string
	Full_name    string
}

type VillagerGetAllResponseModel struct {
	Villager_id   string `json:"villager_id"`
	Villager_code string `json:"villager_code"`
	Home_address  string `json:"home_address"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Tel_number    string `json:"tel_number"`
	Remark        string `json:"remark"`
	Status        string `json:"status"`
	Create_by     string `json:"create_by"`
	Create_data   string `json:"create_date"`
}
