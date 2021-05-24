package constants

//-----------------Database
const DbHost = "cit.bannayuu.com"
const DbName = "uat_cit_bannayuu_db"
const DbPort = "5432"

//-----------------Authen
const jwtAccessToken = "f56c3775-07b0-45e7-800f-304274533cb7"
const mainHTTPClient = "bannayuu/admin/api/v1/"
const authHTTPClient = "authen"

func GetAuthenHTTPClient() string {
	return mainHTTPClient + authHTTPClient
}

func GetJwtAccessToken() string {
	return jwtAccessToken
}
