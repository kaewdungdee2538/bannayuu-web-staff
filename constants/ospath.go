package constants
import (
	"os"
	"fmt"
	"time"
)
var errLogFile *os.File;
var accessLogFile *os.File;
var errLogLoginFile *os.File;
var accessLogLoginFile *os.File;
func SetupOSPath(){
	runningDir, _ := os.Getwd()
	runningDir = fmt.Sprintf("%s/log",runningDir);
	if _, err := os.Stat(runningDir); os.IsNotExist(err) {
		os.Mkdir(runningDir, 0700)
	}
	errLog, _ := os.OpenFile(fmt.Sprintf("%s/api_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accessLog, _ := os.OpenFile(fmt.Sprintf("%s/api_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	errLogLogin, _ := os.OpenFile(fmt.Sprintf("%s/api_login_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accessLogLogin, _ := os.OpenFile(fmt.Sprintf("%s/api_login_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	errLogFile = errLog;
	accessLogFile = accessLog;
	errLogLoginFile = errLogLogin;
	accessLogLoginFile = accessLogLogin;
}
func WriteLog(log *os.File,text string){
	dt := time.Now()
	log.WriteString(fmt.Sprintf("<------------------------------------------------------->\n[Time : %s]\n%s\n",
	dt.Format(time.UnixDate),text))
}
func GetErrorLogFile() *os.File{
	return errLogFile;
}
func GetAccessLogFile() *os.File{
	return accessLogFile;
}
func GetErrorLogLoginFile() *os.File{
	return errLogLoginFile;
}
func GetAccessLogLoginFile() *os.File{
	return accessLogLoginFile;
}