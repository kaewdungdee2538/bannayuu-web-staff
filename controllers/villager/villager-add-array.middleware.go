package controller

import (
	constants "bannayuu-web-admin/constants"
	villager_model "bannayuu-web-admin/model/villager"
	"fmt"
	"reflect"
	"strings"

	"bannayuu-web-admin/utils"
	"github.com/kaewdungdee2538/ouanfunction/validation"
)

func addVillagerMiddleware(villagerModelReq villager_model.VillagerRequestModel) (bool, string) {
	Home_address := strings.ReplaceAll(villagerModelReq.Home_address, " ", "")
	First_name := strings.TrimSpace(villagerModelReq.First_name)
	Last_name := strings.TrimSpace(villagerModelReq.Last_name)
	tel_number := RemoveWhiteSpace(strings.TrimSpace(villagerModelReq.Tel_number))
	if len(Home_address) == 0 {
		return true, constants.MessageHomeAddressNotFound
	} else if validation.IsNotStringAlphabetRemark(Home_address) {
		return true, fmt.Sprintf("%s ,[ที่อยู่ : %s]", constants.MessageHomeAddressProhibitSpecial, Home_address)
	} else if len(First_name) == 0 {
		return true, constants.MessageFirstNameNotFound
	} else if validation.IsNotStringAlphabetForJSONString(First_name) {
		return true, fmt.Sprintf("%s ,[ชื่อ : %s]", constants.MessageFirstNameProhitbitSpecial, First_name)
	} else if len(Last_name) == 0 {
		return true, constants.MessageLastNameNotFound
	} else if validation.IsNotStringAlphabetForJSONString(Last_name) {
		return true, fmt.Sprintf("%s ,[นามสกุล : %s]", constants.MessageLastNameProhitbitSpecial, Last_name)
	} else if len(tel_number) == 0 {
		return true, constants.MessageTelNumberNotFound
	} else if utils.IsNotStringNumber(tel_number) {
		return true, fmt.Sprintf("%s ,[เบอร์โทรศัพท์:%v] %v", constants.MessageTelNumberNotNumber, tel_number, reflect.TypeOf(tel_number))
	} else if validation.IsNotStringAlphabetRemark(villagerModelReq.Remark) {
		return true, constants.MessageRemarkProhibitSpecial
	}else if len(tel_number) < 10 {
		return true,  fmt.Sprintf("%s ,[เบอร์โทรศัพท์ :%s]",constants.MessageTelNumberIsLess10Digit,tel_number)
	} else if len(tel_number) > 10 {
		return true, fmt.Sprintf("%s ,[เบอร์โทรศัพท์ :%s]",constants.MessageTelNumberIsMoreThan10Digit,tel_number)
	}
	return false, ""
}

func RemoveWhiteSpace(input string) string {
	cleanedString := strings.Replace(input, "\u200B", "", -1)
	return cleanedString
}
