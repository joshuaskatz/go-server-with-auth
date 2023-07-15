package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	JWTSigningKey string
	DbHost string
	DbPort int
	DbUser string
	DbName string
	DbPassword string
}

func LoadEnv() Env {
	appEnv := os.Getenv("APP_ENV")

	filePath := GetEnvPath(appEnv)

	err := godotenv.Load(filePath)

	if err != nil {
		panic(err)	
	}
  
	jwtSigningKey := os.Getenv("JWT_SIGNING_KEY")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	parsedDbPort, err := strconv.Atoi(dbPort)
	
    if err != nil {
        panic(err)
    }

	env := Env{
		JWTSigningKey: jwtSigningKey,
		DbHost: dbHost,
		DbPort: parsedDbPort,
		DbUser: dbUser,
		DbName: dbName,
		DbPassword: dbPassword,
	}

	fmt.Println(env)

	return env
}

func GetEnvPath(appEnv string) (string){
	path := ""; 

	if appEnv == "dev" {
		path = ".env.dev"
	}

    filePath, _ := filepath.Abs(path)
	
	return filePath 

}
