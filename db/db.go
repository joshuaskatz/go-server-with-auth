package db

import (
	"database/sql"
	"fmt"
	"server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "password"
	dbname = "postgres"
)

// OpenConnection - opens a connection to the database
func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}

// Init - initialize Database and return a router
func Init()(*gin.Engine){
	db := OpenConnection()
	
	defer db.Close()

	modelsInit(db)

	r := gin.Default()

	return r
}


func modelsInit(db *sql.DB){
	paths := []string{
		"./schema/album/create.sql",
		"./schema/user/create.sql",
	}

	for _, path := range(paths) {
		models.CreateTable(db, path)
	}
}