package appointment

import (
	"github.com/Nopalev/janjiyan/utilities/database"
)

func createDB(appointment Appointment) Appointment {
	db := database.GetDB()
	db.Create(&appointment)
	return appointment
}

func readDB(ID int) Appointment {
	db := database.GetDB()
	var appointment Appointment
	appointment.ID = ID
	db.First(&appointment)
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
	db.Model(&appointment).Where("id = ?", appointment.ID).Updates(appointment)
}

func deleteDB(ID int) {
	db := database.GetDB()
	var appointment Appointment
	appointment.ID = ID
	db.Delete(&appointment)
}
