package constants

import (
	"fmt"
	"os"
)

// -----------------Database
var DbHost = "uat.bannayuu.com"
var DbName = "uat_cit_bannayuu_db"
var DbPort = "5432"
var AppPort = ":4501"

// ----------------demo
// const DbUserName = "postgres"
// const DbPassword = "P@ssw0rd"
// ----------------uat and production
const DbUserName = "cit"
const DbPassword = "db13apr"

var RootImages = "F:/API/myvilla/web-admin/back/uploads/images"
var WEB_MANAGEMENT_RESET_USER = ""

var DbMaxIdleTime = "10"
var DbMaxConnections = "100"

// -----------------Constanst
const CitCompany = "999"
const EmployeeTypeOfManagement = "MANAGEMENT"
const EmployeeTypeOfUser = "USER"
const PrivilegeOfUserTypeId = "5"

const VAR_COMPANY = "company"
const VAR_EMPLOYEE = "employee"
const VAR_UUID = "uuid"

// -----------------Authen
var jwtAccessToken = "f56c3775-07b0-45e7-800f-304274533cb7"

// ----------------Constanst uri
const mainHTTPClient = "bannayuu/admin/api/v1"
const authHTTPClient = "/authen"
const companyHTTPClient = "/company"
const homeHTTPClient = "/home"
const villagerHTTPClient = "/villager"
const userHTTPClient = "/user"
const slotHTTPClient = "/slot"

func InitializeEnv() bool {
	DbHost = os.Getenv("DB_HOST")
	DbName = os.Getenv("DB_NAME")
	DbPort = os.Getenv("DB_PORT")
	AppPort = os.Getenv("APP_PORT")
	RootImages = os.Getenv("ROOT_IMAGE")
	jwtAccessToken = os.Getenv("AUTHEN_TOKEN")
	WEB_MANAGEMENT_RESET_USER = os.Getenv("WEB_MANAGEMENT_RESET_USER")

	DbMaxIdleTime = os.Getenv("DB_MAX_IDLE_TIME")
	DbMaxConnections = os.Getenv("DB_MAX_CONECTIOS")

	fmt.Printf("DbHost : %s\n", DbHost)
	fmt.Printf("DbName : %s\n", DbName)
	fmt.Printf("DbPort : %s\n", DbPort)
	fmt.Printf("AppPort : %s\n", AppPort)
	fmt.Printf("RootImages : %s\n", RootImages)
	fmt.Printf("jwtAccessToken : %s\n", jwtAccessToken)
	fmt.Printf("WEB_MANAGEMENT_RESET_USER : %s\n", WEB_MANAGEMENT_RESET_USER)

	fmt.Printf("DbMaxIdleTime : %s\n", DbMaxIdleTime)
	fmt.Printf("DbMaxConnections : %s\n", DbMaxConnections)
	if DbHost == "" || DbName == "" || DbPort == "" || AppPort == "" || RootImages == "" || jwtAccessToken == "" || WEB_MANAGEMENT_RESET_USER == "" {
		return false
	}
	
	return true
}
func GetHTTPClient() string {
	return mainHTTPClient
}

func GetAuthenHTTPClient() string {
	return mainHTTPClient + authHTTPClient
}
func GetCompanyInsertHTTPClient() string {
	return mainHTTPClient + companyHTTPClient
}

func GetJwtAccessToken() string {
	return jwtAccessToken
}

func GetHomeHTTPClient() string {
	return mainHTTPClient + homeHTTPClient
}

func GetVillagerHTTPClient() string {
	return mainHTTPClient + villagerHTTPClient
}

func GetUserHTTPClient() string {
	return mainHTTPClient + userHTTPClient
}

func GetSlotHTTPClient() string {
	return mainHTTPClient + slotHTTPClient
}
