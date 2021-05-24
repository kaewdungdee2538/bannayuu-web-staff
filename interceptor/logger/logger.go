package logger

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	constants "bannayuu-web-admin/constants"
)

func CreateOrAppendLogger(router *gin.Engine) {
	gin.DefaultErrorWriter = constants.GetErrorLogFile();
	gin.DefaultWriter = constants.GetAccessLogFile();
	// r.Use(gin.Logger())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		//custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n<-------------------------------------------------------------------------->\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))
}
