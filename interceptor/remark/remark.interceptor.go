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
type RemarkModelRequest struct {
	Remark        string
}
type RemarkModelResponse struct {
	Remark        string `json:"remark"`
}
func GetIdCompanyValidateValuesInterceptor(c *gin.Context) {
	buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)
	var remarkModel RemarkModelRequest
	err := json.Unmarshal([]byte(jsonString), &remarkModel)
	if err != nil {
		//--------create error log
		format_utls.WriteLog(format_utls.GetErrorLogFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
		fmt.Printf("Error parsing JSON string - %s", err)
	}
	isErr, msg := checkValuesGetRemark(remarkModel)
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

func checkValuesGetRemark(remarkModel RemarkModelRequest) (bool, string) {
	remark := strings.TrimSpace(remarkModel.Remark)
	if len(remark) == 0 {
		return true, constants.MessageRemarkNotFount;
	}else if len(remark) < 10{ 
		return true, constants.MessageRemarkIsLower10Character;
	}else if format_utls.IsNotStringAlphabetRemark(remark){
		return true, constants.MessageRemarkProhibitSpecial;
	}
	return false, ""
}