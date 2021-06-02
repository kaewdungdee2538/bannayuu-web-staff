package interceptor

import (
	constants "bannayuu-web-admin/constants"
	format_utls "bannayuu-web-admin/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RemarkModelRequest struct {
	Remark string `form:"remark" binding:"required"`
}

func CheckRemarkValidateValueFormDataInterceptor(c *gin.Context) {
	var remarkModel RemarkModelRequest
	if err := c.ShouldBind(&remarkModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": constants.MessageDataNotCompletely})
		c.Abort()
		return
	}
	fmt.Print(remarkModel)
	isErr, msg := checkValuesGetRemark(remarkModel)
	if isErr {
		c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": msg})
		c.Abort()
	} else {
		c.Next()
	}
}

func checkValuesGetRemark(remarkModel RemarkModelRequest) (bool, string) {
	remark := strings.TrimSpace(remarkModel.Remark)
	if len(remark) == 0 {
		return true, constants.MessageRemarkNotFount
	} else if len(remark) < 10 {
		return true, constants.MessageRemarkIsLower10Character
	} else if format_utls.IsNotStringAlphabetRemark(remark) {
		return true, constants.MessageRemarkProhibitSpecial
	}
	return false, ""
}
