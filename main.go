package main

import (
	"log"

	"github.com/Nopalev/janjiyan/routes"
	"github.com/Nopalev/janjiyan/utilities/auth"
	"github.com/Nopalev/janjiyan/utilities/database"
	"github.com/Nopalev/janjiyan/utilities/errorlog"
	"github.com/Nopalev/janjiyan/utilities/migration"
	"github.com/joho/godotenv"
)

func main() {
	auth.Init()
	errorlog.Init()
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	database.Init()

	migration.Migrate()

	r := routes.Routes()
	r.Run(":8080")
}
