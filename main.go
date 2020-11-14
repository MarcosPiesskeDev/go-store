package main

import (
	"go-store-back/src/config"
	"go-store-back/src/routes"
)

func main() {

	routes.ConfigRoutes()
	db, _ := config.GetDb()
	defer db.Close()
}
