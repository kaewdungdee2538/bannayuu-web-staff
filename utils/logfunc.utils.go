package utils

import (
	constants "bannayuu-web-admin/constants"
	"fmt"
	"os"
	"time"
)

func WriteLog(log *os.File, text string) {
	dt := time.Now()
	log.WriteString(fmt.Sprintf("<------------------------------------------------------->\n[Time : %s]\n%s\n",
		dt.Format(time.UnixDate), text))

}
func WriteLogInterface(log *os.File, items map[string]interface{}, text string) {
	dt := time.Now()
	log.WriteString(fmt.Sprintf("<------------------------------------------------------->\n[Time : %s]\n%s\nRequest : %s\n",
		dt.Format(time.UnixDate), text, items))

}
func GetErrorLogFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootMain, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_error.log", directory)
	errLogFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return errLogFile
}
func GetAccessLogFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootMain, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_access.log", directory)
	accessLogFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return accessLogFile
}
func GetErrorLogLoginFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootLogin, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_login_error.log", directory)
	errLogLoginFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return errLogLoginFile
}
func GetAccessLogLoginFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootLogin, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_login_access.log", directory)
	accessLogLoginFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return accessLogLoginFile
}

func GetErrorLogCompanyFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootCompany, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_company_error.log", directory)
	errLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return errLogCompanyFile
}

func GetAccessLogCompanyFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootCompany, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_company_access.log", directory)
	accessLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return accessLogCompanyFile
}

func GetErrorLogHomeFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootHome, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_home_error.log", directory)
	errLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return errLogCompanyFile
}

func GetAccessLogHomeFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootHome, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_home_access.log", directory)
	accessLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return accessLogCompanyFile
}

func GetErrorLogVillagerFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootVillager, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_villager_error.log", directory)
	errLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return errLogCompanyFile
}

func GetAccessLogVillagerFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootVillager, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_villager_access.log", directory)
	accessLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return accessLogCompanyFile
}

func GetErrorLogUserFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootUser, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_user_error.log", directory)
	errLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return errLogCompanyFile
}

func GetAccessLogUserFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootUser, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_user_access.log", directory)
	accessLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return accessLogCompanyFile
}

func GetErrorLogUserResetPasswordFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootUserResetPassword, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_user_reset_password_error.log", directory)
	errLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return errLogCompanyFile
}

func GetAccessLogUserResetPasswordFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootUserResetPassword, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_user_reset_password_access.log", directory)
	accessLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return accessLogCompanyFile
}

func GetErrorLogAddSlotManualFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootSlot, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_add_slot_maunal_error.log", directory)
	errLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return errLogCompanyFile
}

func GetAccessLogAddSlotManualFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootSlot, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_add_slot_maunal_access.log", directory)
	accessLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return accessLogCompanyFile
}


func GetErrorLogGetSlotFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootSlot, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_get_slot_error.log", directory)
	errLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return errLogCompanyFile
}

func GetAccessLogGetSlotFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s/%s", constants.RootSlot, directory_date)
	CheckDirectory(directory)
	root_str := fmt.Sprintf("%s/api_get_slot_access.log", directory)
	accessLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return accessLogCompanyFile
}
