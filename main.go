package main

import (
	"log"
	"os"
	"slices"

	"github.com/Nopalev/janjiyan/routes"
	"github.com/Nopalev/janjiyan/utilities/auth"
	"github.com/Nopalev/janjiyan/utilities/database"
	"github.com/Nopalev/janjiyan/utilities/errorlog"
	"github.com/Nopalev/janjiyan/utilities/migration"
	"github.com/Nopalev/janjiyan/utilities/seeder"
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

	args := os.Args

	if slices.Contains(args, "-s") {
		seeder.Seeder()
	}

	r := routes.Routes()
	r.Run(":8080")
}
