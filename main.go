package main

import (
	"cdn.go/db"
	"cdn.go/routes"
	"github.com/joho/godotenv"
)



func main() {
	err := godotenv.Load()

	if err != nil {
		panic("failed to load env")
	}
	
	db.InitDB()
	routes.InitRoutes()
}