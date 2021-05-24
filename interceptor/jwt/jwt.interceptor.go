package interceptor

import (
	constants "bannayuu-web-admin/constants"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtVerify(c *gin.Context) {
	authenHeader := c.GetHeader("Authorization")
	fmt.Println("jwt verify")
	if authenHeader == "" {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"error": true,
			"result":  nil,
			"message": constants.MessageAuthorizationBearerNotSet})
		c.Abort()
		fmt.Println("Header not have a Authorization")
	} else {
		claims := jwt.MapClaims{}
		tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
		token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected singing method: %v", token.Header["alg"])
			}
			return []byte(constants.GetJwtAccessToken()), nil
		})
		claimss, _ := token.Claims.(jwt.MapClaims)
		if token.Valid {
			fmt.Println("Authorization is valid.")
			c.Set(("jwt_employee_id"), claimss["employee_id"])
			c.Next()
		} else {
			fmt.Println("Authorization is not valid or expire.")
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"error": true,
				"result":  nil,
				"message": constants.MessageNotAuthorization})
			c.Abort()
		}
	}
}
