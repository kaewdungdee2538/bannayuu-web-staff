package interceptor

import (
	constants "bannayuu-web-admin/constants"
	format_utls "bannayuu-web-admin/utils"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"io/ioutil"
	"net/http"
	"strings"
)

type authenUserLoginStc struct {
	Username string
	Password string
}

func LoginValidateValues(c *gin.Context) {
	buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)
	var userModel authenUserLoginStc
	err := json.Unmarshal([]byte(jsonString), &userModel)
	if err != nil {
		//--------create error log
		constants.WriteLog(constants.GetErrorLogFile(),fmt.Sprintf("Error parsing JSON string - %s", err))
		fmt.Printf("Error parsing JSON string - %s", err)
	}
	isErr,msg := checkValues(userModel);
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		//-----forward request body middleware to endpoint
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr2
		c.Next()
	}
}

func checkValues(userModel authenUserLoginStc) (bool,string){
	username := strings.TrimSpace(userModel.Username);
	password := strings.TrimSpace(userModel.Password);
	if len(username) == 0 {
		return true, constants.MessageUsernameNotFount;
	} else if len(password) == 0 {
		 return true,constants.MessagePasswordNotFount;
	} else if format_utls.IsNotStringAlphabet(username){
		return true,constants.MessageUsernameIsSpecialProhibit;
	}else if format_utls.IsNotStringAlphabet(password){
		return true,constants.MessagePasswordIsSpecialProhibit;
	}
	return false,"";
}
