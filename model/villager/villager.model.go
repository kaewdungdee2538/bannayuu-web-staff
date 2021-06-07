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
