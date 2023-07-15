package middleware

import (
	"net/http"
	"server/utils"

	"github.com/gin-gonic/gin"
)
func JWTTokenCheck(c *gin.Context) {
	 err := utils.TokenValid(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Next()
}