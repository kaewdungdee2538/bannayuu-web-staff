package utils

import (
	"encoding/json"
	"fmt"
)

func ConvertInterfaceToJSON(obj map[string]interface{}) (bool, string) {
	obj_result, err := json.Marshal(obj)
	if err != nil {
		//--------create error log
		WriteLog(GetErrorLogLoginFile(), fmt.Sprintf("Error parsing JSON string - %s", err))
		fmt.Printf("Error parsing JSON string - %s", err)
		return true, ""
	}
	return false, string(obj_result)
}
