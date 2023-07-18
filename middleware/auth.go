package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
	"server/config"
	"server/db"
	"server/errors"
	"server/models"
	"server/utils"

	"gopkg.in/gomail.v2"

	"github.com/gin-gonic/gin"
)

var env = config.LoadEnv()

func EmailValidation(c *gin.Context) {

	input, _ := ioutil.ReadAll(c.Request.Body)

	if err := c.BindJSON(&input); err != nil {
		errors.BadRequest(c)
		return
	}

	match, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, input.Email)

	if !match {
		errors.InvalidEmail(c)
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Next()
}

func PasswordValidation(c *gin.Context) {
	password := c.GetString("password")

	// Change this to match your parameters
	match, err := regexp.MatchString(`^.{8,}$`, password)

	if !match {
		errors.InvalidPassword(c)
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Next()
}

func IsVerified(c *gin.Context) {
	var input models.AuthInput

	if err := c.BindJSON(&input); err != nil {
		errors.BadRequest(c)
		return
	}

	if !input.Verified {
		errors.UserNotVerified(c)
		return
	}

	c.Next()
}

func VerificationEmail(c *gin.Context) {
	DB := db.OpenConnection()

	defer DB.Close()

	var input models.EmailTemplate

	if err := c.BindJSON(&input); err != nil {
		errors.BadRequest(c)
		return
	}

	queryFilePath, _ := filepath.Abs("./schema/user/insert.sql")

	query := utils.ParseFile(queryFilePath)

	jwt, err := utils.GenerateJWT(input.Email)

	if err != nil {
		errors.BadRequest(c)
		return
	}

	sqlStatement := fmt.Sprintf(query, jwt, input.Email)

	if _, err := DB.Query(sqlStatement); err != nil {
		errors.BadRequest(c)
		return
	}

	url := env.ServerUrl + "/verify/" + jwt

	htmlFilePath, _ := filepath.Abs("./templates/email.html")

	html := utils.ParseFile(htmlFilePath)

	body := fmt.Sprintf(html, input.Email, url)

	m := gomail.NewMessage()
	m.SetHeader("From", "joshua.samuel.katz@gmail.com")
	m.SetHeader("To", input.Email)
	m.SetHeader("Subject", "Verify your email!")
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer("smtp.example.com", 587, "user", "123456")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	c.Next()
}
