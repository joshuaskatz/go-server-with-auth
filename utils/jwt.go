package utils

import (
	"fmt"
	"server/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// TODO(@joshuaskatz: add env vars)
var env = config.LoadEnv()

func GenerateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["authorized"] = true
	claims["user"] = email
	tokenString, err := token.SignedString(env.JWTSigningKey)

	println(tokenString, err)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//   This takes a JWT token as input and extracts the key used for verifying the token's signature.
//   It checks if the token's signing method is HMAC, and if not, it returns an error indicating an unexpected signing method.
//   If the signing method is HMAC, it returns the key as a byte array and a nil error.
func GetSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(env.JWTSigningKey), nil
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)

	_, err := jwt.Parse(tokenString, GetSecretKey)

	if err != nil {
		return err
	}

	return nil
}

func ExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

