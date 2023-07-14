package main

import (
	"server/db"
	"server/routes"
)





func main(){
	router := db.Init()

	// Routes
	routes.AlbumRoutes(router)
	

	router.Run("localhost:8080")
}


