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
