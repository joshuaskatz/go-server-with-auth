package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(r *gin.RouterGroup) {
	r.Use(middleware.JWTTokenCheck).GET("/albums", controllers.GetAlbums)
	r.Use(middleware.JWTTokenCheck).POST("/albums", controllers.PostAlbum)
}
