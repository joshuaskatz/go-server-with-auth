package controllers

import (
	"fmt"
	"path/filepath"
	"server/db"
	"server/models"
	"server/schema"
)

func getUser (email string) (models.User, error) {
	DB := db.OpenConnection()

	defer DB.Close()
	
	filePath, _ := filepath.Abs("./schema/user/select.sql")

    query := schema.ParseFile(filePath)	
	
	sqlStatement := fmt.Sprintf(query, email)

	row := DB.QueryRow(sqlStatement)

	err := row.Err()

	if err != nil {
		return models.User{}, err
	}

	var user models.User


	if err := row.Scan(&user.ID, &user.Role, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return models.User{}, err
	}
	
	return user, nil
}

func validatePasswordHash(u *models.User, passwordHash string) bool {
	return u.PasswordHash == passwordHash
}

func createUser(email string, passwordHash string, role string, name string) (models.User, error) {
	DB := db.OpenConnection()

	defer DB.Close()

	filePath, _ := filepath.Abs("./schema/album/insert.sql")

    query :=  schema.ParseFile(filePath)

	sqlStatement := fmt.Sprintf(query, email, passwordHash, name, role)

	res, err := DB.Query(sqlStatement)

	if err != nil {
		 return models.User{}, err
	}

	var user models.User


	if err := res.Scan(&user.CreatedAt, &user.Email, &user.ID, &user.Name, &user.PasswordHash, &user.Role, &user.UpdatedAt); err != nil {
		return models.User{}, err
    }

	return user, nil
}