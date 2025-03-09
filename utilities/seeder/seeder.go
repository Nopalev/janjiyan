package seeder

import (
	"time"

	"github.com/Nopalev/janjiyan/domains/appointment"
	"github.com/Nopalev/janjiyan/domains/user"
)

func Seeder() {
	_, offset := time.Now().Local().Zone()
	user.Register(
		user.User{
			Name:     "testName",
			Username: "user_1",
			Timezone: offset,
			Password: "p4ssw0rd"})
	_, offset = time.Now().UTC().Zone()
	user.Register(
		user.User{
			Name:     "testName",
			Username: "user_2",
			Timezone: offset,
			Password: "p4ssw0rd"})

	appointment.Create(
		appointment.Appointment{
			Title:     "first",
			Start:     time.Now(),
			End:       time.Now().Add(time.Hour),
			CreatorID: 1},
		"user_1")
	appointment.Create(
		appointment.Appointment{
			Title:     "second",
			Start:     time.Now().Add(24 * time.Hour),
			End:       time.Now().Add(25 * time.Hour),
			CreatorID: 1},
		"user_1")
}
