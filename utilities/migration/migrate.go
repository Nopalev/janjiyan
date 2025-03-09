package migration

import (
	"github.com/Nopalev/janjiyan/domains/appointment"
	"github.com/Nopalev/janjiyan/domains/invitation"
	"github.com/Nopalev/janjiyan/domains/user"
	"github.com/Nopalev/janjiyan/utilities/database"
)

func Migrate() {
	db := database.GetDB()
	if db == nil {
		return
	}

	db.Debug().AutoMigrate(&user.User{}, &appointment.Appointment{}, &invitation.Invitation{})
}
