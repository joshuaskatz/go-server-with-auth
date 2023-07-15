package main

import (
	"server/db"
	"server/routes"
)





func main(){
	router := db.Init()

	public := router.Group("/api")
	protected := router.Group("/api/admin")

	// Routes
	routes.AlbumRoutes(protected)
	routes.AuthRoutes(public)
	 

	err := router.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}


