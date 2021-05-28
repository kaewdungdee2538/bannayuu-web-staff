package constants

import (
	"fmt"
	"os"
)

var RootMain string
var RootLogin string
var RootCompany string

func SetupOSPath() {
	runningDir, _ := os.Getwd()
	runningDir = fmt.Sprintf("%s\\log", runningDir)
	//---------------------------------------------------------
	RootMain = fmt.Sprintf("%s\\main", runningDir)
	RootLogin = fmt.Sprintf("%s\\login", runningDir)
	RootCompany = fmt.Sprintf("%s\\company", runningDir)
	//----------------------------------------------------------//
}
