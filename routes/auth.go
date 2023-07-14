package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(r *gin.Engine){
	r.GET("/albums", controllers.GetAlbums)
	r.POST("/albums", controllers.PostAlbum)
}