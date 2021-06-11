package interceptor

import (
	constants "bannayuu-web-admin/constants"
	intercep_home "bannayuu-web-admin/interceptor/home"
	format_utls "bannayuu-web-admin/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type VillagerGetAllInterceptorModel struct {
	Company_id   string
	Home_address string
	Full_name    string
}

func CheckGetVillagerValueInterceptor(c *gin.Context) {
	var villagerModel VillagerGetAllInterceptorModel
	buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)

	err := json.Unmarshal([]byte(jsonString), &villagerModel)

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
	fmt.Print(villagerModel)
	isErr, msg := checkValueFullName(villagerModel)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		// ---------Convert obj to json string
		homeInfo, err := json.Marshal(villagerModel)
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
func checkValueFullName(villagerModel VillagerGetAllInterceptorModel) (bool, string){
	if format_utls.IsNotStringAlphabetRemark(villagerModel.Full_name){
		return true, constants.MessageFullNameProhitbitSpecial
	}
	return checkValueHomeAddress(villagerModel)
}
func checkValueHomeAddress(villagerModel VillagerGetAllInterceptorModel) (bool, string) {
	if format_utls.IsNotStringAlphabetRemark(villagerModel.Home_address) {
		return true, constants.MessageHomeAddressProhibitSpecial
	}
	return intercep_home.CheckValueCompanyIdNotDisavle(villagerModel.Company_id)
}
