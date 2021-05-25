package api

import (
	constants "bannayuu-web-admin/constants"
	db "bannayuu-web-admin/db"
	model_company "bannayuu-web-admin/model/company"

	"fmt"
	"net/http"
	// "io/ioutil"
	"github.com/gin-gonic/gin"
)

func AddCompany(c *gin.Context) {
	// buf, _ := ioutil.ReadAll(c.Request.Body) // handle the error
	// jsonString := string(buf)
	jwtemployeeid,_ := c.Get("jwt_employee_id");
	fmt.Printf("companyModel : %v ", jwtemployeeid)
	var companyModelReq model_company.CompanyAddModelRequest
	if c.ShouldBind(&companyModelReq) == nil {
		fmt.Printf("companyModel : %v ", companyModelReq)
		query := `
		insert into m_company(
			company_code
			,company_name
			,company_promotion
			,company_start_date
			,company_expire_date
			,create_by
			,create_date
		) values(
			@company_code
			,@company_name
			,@company_promotion
			,@company_start_date
			,@company_expire_date
			,@create_by
			,current_timestamp
			);
			`
		values := map[string]interface{}{
			"company_code":        companyModelReq.Company_code,
			"company_name":        companyModelReq.Company_name,
			"company_promotion":   companyModelReq.Company_promotion,
			"company_start_date":  companyModelReq.Company_start_date,
			"company_expire_date": companyModelReq.Company_expire_date,
			"create_by":           jwtemployeeid}

		if err, message := db.SaveTransactionDB(query, values); err {
			fmt.Printf("add company error : %s", message)
			c.JSON(http.StatusOK, gin.H{"error": true, "result": nil, "message": message})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": false, "result": nil, "message": constants.MessageSuccess})
		}
	}
}
