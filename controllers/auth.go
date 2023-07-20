package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"server/db"
	"server/errors"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Verify(c *gin.Context) {
	DB := db.OpenConnection()

	defer DB.Close()

	jwt := c.Param("jwt")
	fmt.Println(jwt)

	claims, err := utils.ExtractClaims(jwt)

	if err != nil {
		errors.BadRequest(c)
		return
	}

	email := claims["user"].(string)

	queryFilePath, _ := filepath.Abs("./schema/user/verify.sql")

	query := utils.ParseFile(queryFilePath)

	sqlStatement := fmt.Sprintf(query, email)

	if _, err := DB.Query(sqlStatement); err != nil {
		errors.BadRequest(c)
		return
	}

	c.Writer.WriteHeader(204)
}

func Login(c *gin.Context) {
	DB := db.OpenConnection()

	defer DB.Close()

	var input models.AuthInput

	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		errors.BadRequest(c)
		return
	}

	user, err := GetUser(input.Email)

	if err != nil {
		errors.UserNotFound(c)
		return
	}

	if match := utils.CheckPasswordHash(input.Password, user.PasswordHash); !match {
		errors.BadRequest(c)
		return
	}

	jwt, jwtErr := utils.GenerateJWT(user.Email, false)

	if jwtErr != nil {
		errors.BadRequest(c)
		return
	}

	c.IndentedJSON(http.StatusCreated, jwt)
}

func Register(c *gin.Context) {
	DB := db.OpenConnection()

	defer DB.Close()

	var input models.AuthInput

	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		errors.BadRequest(c)
		return
	}

	filePath, _ := filepath.Abs("./schema/user/insert.sql")

	query := utils.ParseFile(filePath)

	passwordHash, hashErr := utils.HashPassword(input.Password)

	if hashErr != nil {
		errors.BadRequest(c)
	}

	sqlStatement := fmt.Sprintf(query, input.Email, passwordHash)

	if _, err := DB.Query(sqlStatement); err != nil {
		errors.EmailInUse(c)
	}

	c.Writer.WriteHeader(204)
}

func GetUser(email string) (models.User, error) {
	DB := db.OpenConnection()

	defer DB.Close()

	filePath, _ := filepath.Abs("./schema/user/select.sql")

	query := utils.ParseFile(filePath)

	sqlStatement := fmt.Sprintf(query, email)

	row := DB.QueryRow(sqlStatement)

	if err := row.Err(); err != nil {
		return models.User{}, err
	}

	var user models.User

	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.Verified, &user.VerificationCode); err != nil {
		return models.User{}, err
	}

	return user, nil
}
