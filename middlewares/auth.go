package middlewares

import (
	"example/goUdemyRest/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "user not authorized"})
		return
	}

	uId, err := utils.VerifyJWTToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "user not authorized"})
		return
	}

	c.Set("userId", uId)
	c.Next()

}
