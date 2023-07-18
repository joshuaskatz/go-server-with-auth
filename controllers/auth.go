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
)

func Login(c *gin.Context) {
	DB := db.OpenConnection()

	defer DB.Close()

	var input models.AuthInput

	if err := c.BindJSON(&input); err != nil {
		errors.BadRequest(c)
		return
	}

	user, err := getUser(input.Email)

	if err != nil {
		errors.UserNotFound(c)
		return
	}

	if match := utils.CheckPasswordHash(input.Password, user.PasswordHash); !match {
		errors.BadRequest(c)
		return
	}

	jwt, jwtErr := utils.GenerateJWT(user.Email)

	println(&jwtErr)

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

	if err := c.BindJSON(&input); err != nil {
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

func getUser(email string) (models.User, error) {
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

	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return models.User{}, err
	}

	return user, nil
}
