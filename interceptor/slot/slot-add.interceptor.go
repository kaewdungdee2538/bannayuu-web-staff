package interceptor

import (
	constants "bannayuu-web-admin/constants"
	home_intercep "bannayuu-web-admin/interceptor/home"
	slot_add_model "bannayuu-web-admin/model/slot/add"
	format_utls "bannayuu-web-admin/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func CheckAddSlotValueInterceptor(c *gin.Context) {
	var request slot_add_model.SlotAddRequest
	buf, _ := io.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)

	err := json.Unmarshal([]byte(jsonString), &request)

	if err != nil {
		//--------create error log
		format_utls.WriteLog(format_utls.GetErrorLogUserFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
		fmt.Printf("Error parsing JSON string - %s", err)
	}

	isErr, msg := checkSlotAddRequest(request)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		// ---------Convert obj to json string
		userInfo, err := json.Marshal(request)
		if err != nil {
			format_utls.WriteLogInterface(format_utls.GetErrorLogUserFile(), nil, constants.MessageCovertObjTOJSONFailed)
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
			return
		}
		// -----forward request body middleware to endpoint
		rdr2 := io.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf("%v", string(userInfo)))))
		c.Request.Body = rdr2
		c.Next()
	}
}
func checkSlotAddRequest(request slot_add_model.SlotAddRequest) (bool, string) {
	company_id := request.Company_id
	slot_count := request.Slot_count
	guardhouse_id := request.Guardhouse_id
	guardhouse_code := request.Guardhouse_code
	
	if slot_count == 0 {
		return true, constants.MessageSlotCountNotFound
	} else if guardhouse_id == 0 {
		return true, constants.MessageGuadhouseIdNotFound
	} else if len(guardhouse_code) == 0 {
		return true, constants.MessageGuadhouseCodeNotFound
	} else if format_utls.IsNotStringEngOrNumber(guardhouse_code) {
		return true, constants.MessageGuadhouseCodeNotEngOrNumber
	} 
	errComId, msgComId := home_intercep.CheckValueCompanyIdNotDisavle(fmt.Sprint(company_id))
	if errComId {
		return true, msgComId
	}
	return false,""
}
