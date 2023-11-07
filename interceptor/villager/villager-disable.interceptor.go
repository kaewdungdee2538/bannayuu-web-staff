package interceptor

import (
	constants "bannayuu-web-admin/constants"
	villager_model "bannayuu-web-admin/model/villager"
	format_utls "bannayuu-web-admin/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaewdungdee2538/ouanfunction/validation"
)

func CheckDisableVillagerValueInterceptor(c *gin.Context) {
	var request villager_model.VillageDisableRequest
	buf, _ := io.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)

	err := json.Unmarshal([]byte(jsonString), &request)

	if err != nil {
		//--------create error log
		format_utls.WriteLog(format_utls.GetErrorLogVillagerDisableFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
		fmt.Printf("Error parsing JSON string - %s", err)
	}

	errValidation := checkSlotDisableRequest(request)
	if errValidation != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": errValidation.Error()})
		c.Abort()
	} else {
		// ---------Convert obj to json string
		userInfo, err := json.Marshal(request)
		if err != nil {
			format_utls.WriteLogInterface(format_utls.GetErrorLogVillagerDisableFile(), nil, constants.MessageCovertObjTOJSONFailed)
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
			return
		}
		// -----forward request body middleware to endpoint
		rdr2 := io.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf("%v", string(userInfo)))))
		c.Request.Body = rdr2
		c.Next()
	}
}
func checkSlotDisableRequest(request villager_model.VillageDisableRequest)  error {
	home_line_id := request.Home_line_id
	remark := request.Remark
	
	if errHomeLine := checkValueHomeLineId(fmt.Sprint(home_line_id));errHomeLine!=nil{
		return errHomeLine
	}else if len(remark) > 0 && validation.IsNotStringAlphabetRemark(remark){
		return errors.New(constants.MessageRemarkProhibitSpecial)
	}
	return nil
}
