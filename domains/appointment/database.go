package appointment

import (
	"github.com/Nopalev/janjiyan/utilities/database"
)

func createDB(appointment *Appointment) {
	db := database.GetDB()
	db.Create(&appointment)
}

func readDB(ID int) Appointment {
	db := database.GetDB()
	var appointment Appointment
	appointment.ID = ID
	db.First(&appointment)
	db.Model(appointment).Association("User").Find(&appointment.User)
	return appointment
}

func readCreatedDB(ID int) []Appointment {
	db := database.GetDB()
	var appointments []Appointment
	db.Where("creator_id = ?", ID).Find(&appointments)
	return appointments
}

func updateDB(appointment Appointment) {
	db := database.GetDB()
	db.Model(&appointment).Where("id = ?", appointment.ID).Update("title", appointment.Title).Update("start", appointment.Start).Update("end", appointment.End)
}

func deleteDB(ID int) {
	db := database.GetDB()
	var appointment Appointment
	appointment.ID = ID
	db.Delete(&appointment)
}

func deleteByUserDB(userID int) {
	db := database.GetDB()
	db.Delete(&Appointment{}, "creator_id = ?", userID)
}
