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
)

type HomeAddressInterceptorModel struct {
	Company_id   string
	Home_address string
}

func CheckHomeAddressValueInterceptor(c *gin.Context) {
	var homeAddressModel HomeAddressInterceptorModel
	buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)

	err := json.Unmarshal([]byte(jsonString), &homeAddressModel)

	if err != nil {
		//--------create error log
		// format_utls.WriteLog(format_utls.GetErrorLogFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
		fmt.Printf("Error parsing JSON string - %s", err)
	}

	// if err := c.ShouldBind(&companyModel); err != nil {
	// 	c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCombineFailed})
	// 	c.Abort()
	// 	return
	// }
	fmt.Print(homeAddressModel)
	isErr, msg := CheckValueHomeAddressNotDisavle(homeAddressModel)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		// ---------Convert obj to json string
		homeInfo, err := json.Marshal(homeAddressModel)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
			return
		}
		// -----forward request body middleware to endpoint
		rdr2 := ioutil.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf("%v", string(homeInfo)))))
		c.Request.Body = rdr2
		c.Next()
	}
}

func CheckValueHomeAddressNotDisavle(homeAddressObj HomeAddressInterceptorModel) (bool, string) {
	if format_utls.IsNotStringAlphabetRemark(homeAddressObj.Home_address) {
		return true, constants.MessageHomeAddressProhibitSpecial
	}
	return CheckValueCompanyIdNotDisavle(homeAddressObj.Company_id)
}
