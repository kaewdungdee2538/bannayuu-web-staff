package api

import (
	constants "bannayuu-web-admin/constants"
	"bannayuu-web-admin/utils"
	db "bannayuu-web-admin/db"
	interceptor "bannayuu-web-admin/interceptor/jwt"
	login_interceptor "bannayuu-web-admin/interceptor/authen"
	model_authen "bannayuu-web-admin/model/authen"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"database/sql"
)

func SetupAuthenAPI(router *gin.Engine) {
	authApiHTTP := constants.GetAuthenHTTPClient()
	fmt.Printf("authen api http : %s", authApiHTTP)
	authenApi := router.Group(authApiHTTP)
	{
		authenApi.POST("/login", login_interceptor.LoginValidateValues,login)
		authenApi.POST("/test", interceptor.JwtVerify)
	}
}

func login(c *gin.Context) {
	var user model_authen.AuthenUserLoginStc
	var userResponseDb model_authen.AuthenUserLoginForDBStc
	if c.ShouldBind(&user) == nil {
		fmt.Printf("username : %s ,password : %s", user.Username, user.Password)
		userName := user.Username
		password := user.Password
		sql1 := `select employee_id, employee_code,first_name_th,last_name_th,username,(passcode = crypt(@password, passcode)) as password_status 
        ,me.company_id,mc.company_name,me.company_list as company_list
		,mep.login_staff_data as privilege_info
         FROM m_employee me 
         inner join m_employee_privilege mep on me.employee_privilege_id = mep.employee_privilege_id
         left join m_company mc on me.company_id = mc.company_id
         WHERE me.username = @username and me.delete_flag = 'N' and mep.delete_flag ='N' and mep.login_staff_status='Y'`
		 rows, _ := db.GetDB().Raw(sql1,sql.Named("password",password),sql.Named("username",userName)).Rows();
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&userResponseDb)
			db.GetDB().ScanRows(rows, &userResponseDb)
			// do something
		}
		if userResponseDb.Employee_id == 0 || !userResponseDb.Password_status {
			//-------convert struct to json
			json_user, err := json.Marshal(user)
			if err != nil {
				fmt.Println(err)
				utils.WriteLog(utils.GetErrorLogLoginFile(),fmt.Sprintf("Login Error : %v\n", err));
				return
			}
			//--------create log
			utils.WriteLog(utils.GetAccessLogLoginFile(),fmt.Sprintf("Username or Password Not Valid !\nRequest : %v\n", string(json_user)))
			//--------response
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageUsernameOrPasswordNotValid})
			return
		}

		//--------------convert json string to object
		var privilege_info map[string]interface{}
		j := userResponseDb.Privilege_info
		err := json.Unmarshal([]byte(j), &privilege_info)
		if err != nil {
			//--------create error log
			utils.WriteLog(utils.GetErrorLogLoginFile(),fmt.Sprintf("Error parsing JSON string - %s", err))
			fmt.Printf("Error parsing JSON string - %s", err)
		}
		//--------------Create JWT Token
		token, _ := createToken(userResponseDb.Employee_id)
		//--------------Response
		newUserResponse := model_authen.AuthenUserLoginResponseStc{
			Access_token:    token,
			Employee_id:     userResponseDb.Employee_id,
			Employee_code:   userResponseDb.Employee_code,
			First_name_th:   userResponseDb.First_name_th,
			Last_name_th:    userResponseDb.Last_name_th,
			Username:        userResponseDb.Username,
			Password_status: userResponseDb.Password_status,
			Company_id:      userResponseDb.Company_id,
			Company_name:    userResponseDb.Company_name,
			Company_list:    userResponseDb.Company_list,
			Privilege_info:  privilege_info,
		}
		//-------convert struct to json
		json_user, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			utils.WriteLog(utils.GetErrorLogLoginFile(),fmt.Sprintf("Login Error : %v\n", err))
			return
		}
		//--------create log
		utils.WriteLog(utils.GetAccessLogLoginFile(),fmt.Sprintf("Login Success.\nRequest : %v\n", string(json_user)))
		//--------response
		c.JSON(http.StatusOK, gin.H{"error": false,
			"result":  newUserResponse,
			"message": "success"})
	} else {
		//--------create log
		utils.WriteLog(utils.GetAccessLogLoginFile(),"map struct error !\n")
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": "map struct error"})
	}
}

func createToken(userId int) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["employee_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken := constants.GetJwtAccessToken()
	os.Setenv("ACCESS_SECRET", accessToken) //this should be in an env file
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
