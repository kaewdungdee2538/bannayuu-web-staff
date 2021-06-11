package constants

//-----------------Database
const DbHost = "uat.bannayuu.com"
const DbName = "uat_cit_bannayuu_db"
const DbPort = "5432"
const AppPort = ":4501"
const CitCompany = "999"
const EmployeeTypeOfManagement = "MANAGEMENT"
const EmployeeTypeOfUser = "USER"
const PrivilegeOfUserTypeId = "5"
const RootImages = "F:\\API\\myvilla\\web-admin\\back\\uploads\\images"

//-----------------Authen
const jwtAccessToken = "f56c3775-07b0-45e7-800f-304274533cb7"
const mainHTTPClient = "bannayuu/admin/api/v1"
const authHTTPClient = "/authen"
const companyHTTPClient = "/company"
const homeHTTPClient= "/home"
const villagerHTTPClient = "/villager"
const userHTTPClient = "/user"

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

func GetHomeHTTPClient() string{
	return mainHTTPClient + homeHTTPClient
}

func GetVillagerHTTPClient() string{
	return mainHTTPClient + villagerHTTPClient
}

func GetUserHTTPClient() string {
	return mainHTTPClient + userHTTPClient
}

