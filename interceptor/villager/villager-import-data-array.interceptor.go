package interceptor

import (
	constants "bannayuu-web-admin/constants"
	villager_model "bannayuu-web-admin/model/villager"
	home_intercep "bannayuu-web-admin/interceptor/home"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)


func CheckAddVillagerArrayValuesInterceptor(c *gin.Context) {
	var villagerObj villager_model.VillagerAddRequestModel
	buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	jsonString := string(buf)

	err := json.Unmarshal([]byte(jsonString), &villagerObj)

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
	fmt.Print(villagerObj)
	isErr, msg := home_intercep.CheckValueCompanyIdNotDisavle(villagerObj.Company_id)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		// ---------Convert obj to json string
		companyInfo, err := json.Marshal(villagerObj)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageCovertObjTOJSONFailed})
			return
		}
		// -----forward request body middleware to endpoint
		rdr2 := ioutil.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf("%v", string(companyInfo)))))
		c.Request.Body = rdr2
		c.Next()
	}
}


