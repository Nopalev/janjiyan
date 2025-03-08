package migration

import (
	"github.com/Nopalev/janjiyan/domains/user"
	"github.com/Nopalev/janjiyan/utilities/database"
)

func Migrate() {
	db := database.GetDB()
	if db == nil {
		return
	}

	db.AutoMigrate(&user.User{})
}
