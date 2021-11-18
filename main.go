package main

import (
	"myproject/mig"
	"myproject/routes"

	"github.com/joho/godotenv"
)

func main() {
	//load .env file
	godotenv.Load(".env")
	// auto migrate
	mig.Migrate()
	// routes
	routes.Routes()
}
