package api
import (
	"github.com/gin-gonic/gin"
	"bannayuu-web-admin/db"
	authen "bannayuu-web-admin/api/authen"
)

func Setup(router *gin.Engine){
	db.SetupDB();
	authen.SetupAuthenAPI(router);
}