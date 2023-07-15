package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"server/db"
	"server/errors"
	"server/models"
	"server/schema"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	DB := db.OpenConnection()

	defer DB.Close()

	filePath, _ := filepath.Abs("./schema/album/select.sql")

	query := schema.ParseFile(filePath)

	rows, err := DB.Query(query)

	if err != nil {
		errors.AlbumsNotFound(c)
		return
	}

	defer rows.Close()

	var albums []models.Album

	for rows.Next() {
		var album models.Album

		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price, &album.CreatedAt, &album.UpdatedAt); err != nil {
			errors.AlbumsNotFound(c)
			return
		}

		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		errors.AlbumsNotFound(c)
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	DB := db.OpenConnection()

	defer DB.Close()

	var input models.Album

	if err := c.BindJSON(&input); err != nil {
		errors.BadRequest(c)
		return
	}

	filePath, _ := filepath.Abs("./schema/album/insert.sql")

	query := schema.ParseFile(filePath)

	sqlStatement := fmt.Sprintf(query, input.Artist, input.Price, input.Title)

	if _, err := DB.Query(sqlStatement); err != nil {
		errors.BadRequest(c)
		return
	}

	// Return 204
	c.Writer.WriteHeader(204)
}
