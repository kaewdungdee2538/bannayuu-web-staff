package utils
import (
	"fmt"
	"os"
	"time"
	constants "bannayuu-web-admin/constants"
)

func WriteLog(log *os.File, text string) {
	dt := time.Now()
	log.WriteString(fmt.Sprintf("<------------------------------------------------------->\n[Time : %s]\n%s\n",
		dt.Format(time.UnixDate), text))
}

func GetErrorLogFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s\\%s", constants.RootMain, directory_date)
	CheckDirectory(directory);
	root_str := fmt.Sprintf("%s\\api_error.log", directory)
	errLogFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return errLogFile
}
func GetAccessLogFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s\\%s", constants.RootMain, directory_date)
	CheckDirectory(directory);
	root_str := fmt.Sprintf("%s\\api_access.log", directory)
	accessLogFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return accessLogFile
}
func GetErrorLogLoginFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s\\%s", constants.RootLogin, directory_date)
	CheckDirectory(directory);
	root_str := fmt.Sprintf("%s\\api_login_error.log", directory)
	errLogLoginFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return errLogLoginFile
}
func GetAccessLogLoginFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s\\%s", constants.RootLogin, directory_date)
	CheckDirectory(directory);
	root_str := fmt.Sprintf("%s\\api_login_access.log", directory)
	accessLogLoginFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return accessLogLoginFile
}

func GetErrorLogCompanyFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s\\%s", constants.RootCompany, directory_date)
	CheckDirectory(directory);
	root_str := fmt.Sprintf("%s\\api_company_error.log", directory)
	errLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return errLogCompanyFile
}

func GetAccessLogCompanyFile() *os.File {
	directory_date := GetDirectoryDate()
	directory := fmt.Sprintf("%s\\%s", constants.RootCompany, directory_date)
	CheckDirectory(directory);
	root_str := fmt.Sprintf("%s\\api_company_access.log", directory)
	accessLogCompanyFile, _ := os.OpenFile(root_str, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	return accessLogCompanyFile
}
