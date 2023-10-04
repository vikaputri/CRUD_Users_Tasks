package main

import (
	"github.com/vikaputri/CRUD_Users_Tasks/routers"

	"github.com/vikaputri/CRUD_Users_Tasks/database"
)

func main() {
	database.StartDB()
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
