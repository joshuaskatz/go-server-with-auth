package main

import (
	"server/config"
	"server/db"
	"server/routes"
)

var env = config.LoadEnv()

func main() {
	router := db.Init()

	public := router.Group("/api")
	protected := router.Group("/api/admin")

	// Routes
	routes.AlbumRoutes(protected)
	routes.AuthRoutes(public)

	if err := router.Run(env.ServerUrl); err != nil {
		panic(err)
	}
}
