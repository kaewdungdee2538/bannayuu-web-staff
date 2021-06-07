package controller

import (
	"strings"
	"bannayuu-web-admin/utils"
	"fmt"
	constants "bannayuu-web-admin/constants"
	villager_model "bannayuu-web-admin/model/villager"
)

func addVillagerMiddleware(villagerModelReq *villager_model.VillagerRequestModel) (bool, string) {
	Home_address := strings.TrimSpace(villagerModelReq.Home_address)
	First_name := strings.TrimSpace(villagerModelReq.First_name)
	Last_name := strings.TrimSpace(villagerModelReq.Last_name)
	Tel_number := strings.TrimSpace(villagerModelReq.Tel_number)
	if len(Home_address) == 0 {
		return true, constants.MessageHomeAddressNotFound
	} else if utils.IsNotStringAlphabetRemark(Home_address) {
		return true, fmt.Sprintf("%s ,[ที่อยู่ : %s]", constants.MessageHomeAddressProhibitSpecial, Home_address)
	} else if len(First_name) == 0 {
		return true, constants.MessageFirstNameNotFound
	} else if utils.IsNotStringAlphabet(First_name) {
		return true,  fmt.Sprintf("%s ,[ชื่อ : %s]",constants.MessageFirstNameProhitbitSpecial,First_name)
	} else if len(Last_name) == 0 {
		return true, constants.MessageLastNameNotFound
	} else if utils.IsNotStringAlphabet(Last_name) {
		return true, fmt.Sprintf("%s ,[นามสกุล : %s]",constants.MessageLastNameProhitbitSpecial,Last_name) 
	} else if len(Tel_number) == 0 {
		return true, constants.MessageTelNumberNotFound
	} else if len(Tel_number) < 10 {
		return true,  fmt.Sprintf("%s ,[เบอร์โทรศัพท์ : %s]",constants.MessageTelNumberIsLess10Digit,Tel_number) 
	} else if len(Tel_number) > 10 {
		return true, fmt.Sprintf("%s ,[เบอร์โทรศัพท์ : %s]",constants.MessageTelNumberIsMoreThan10Digit,Tel_number) 
	} else if utils.IsNotStringNumber(Tel_number) {
		return true,  fmt.Sprintf("%s ,[เบอร์โทรศัพท์ : %s]",constants.MessageTelNumberNotNumber,Tel_number)
	} else if utils.IsNotStringAlphabetRemark(villagerModelReq.Remark) {
		return true, constants.MessageRemarkProhibitSpecial
	}
	return false, ""
}