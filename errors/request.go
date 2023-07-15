package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AlbumsNotFound(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "albums not found"})

}

func UserNotFound(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func EmailInUse(c *gin.Context) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "email is in use"})
}

func BadRequest(c *gin.Context) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"message": "there was an issue with your request",
	})
}
