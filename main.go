package main

import (
	"server/db"
	"server/routes"
)

func main() {
	router := db.Init()

	public := router.Group("/api")
	protected := router.Group("/api/admin")

	// Routes
	routes.AlbumRoutes(protected)
	routes.AuthRoutes(public)

	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
