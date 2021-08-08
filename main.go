package main

import (
	controllers "jokibro/app/http/controller"
	"jokibro/config/database"
	"jokibro/routers"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	godotenv.Load()

	// initDB
	db := database.InitDB()

	// init route
	c := controllers.Controller{
		DB: db,
	}
	mainRoute := routers.MainRoute{
		Controller: &c,
	}

	// migrate table
	database.MigrateTables(db)

	mainRoute.RegisterRoute()
}
