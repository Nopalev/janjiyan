package invitation

import (
	"github.com/Nopalev/janjiyan/utilities/database"
)

func createDB(invitation *Invitation) {
	db := database.GetDB()
	db.Create(&invitation)
}

func readDB(ID int) Invitation {
	db := database.GetDB()
	var invitation Invitation
	invitation.ID = ID
	db.First(&invitation)
	return invitation
}

func assocUser(invitation *Invitation) {
	db := database.GetDB()
	db.Model(invitation).Association("User").Find(&invitation.User)
}

func assocAppointment(invitation *Invitation) {
	db := database.GetDB()
	db.Model(invitation).Association("Appointment").Find(&invitation.Appointment)
	db.Model(invitation.Appointment).Association("User").Find(&invitation.Appointment.User)
}

func readByAppointmentDB(appointmentID int) []Invitation {
	db := database.GetDB()
	var invitations []Invitation
	db.Where("appointment_id = ?", appointmentID).Find(&invitations)
	return invitations
}

func readByUserDB(userID int) []Invitation {
	db := database.GetDB()
	var invitations []Invitation
	db.Where("invitee_id = ?", userID).Find(&invitations)
	return invitations
}

func readByAppointmentsDB(appointmentsID []int) []Invitation {
	db := database.GetDB()
	var invitations []Invitation
	db.Where("appointment_id IN ?", appointmentsID).Find(&invitations)
	return invitations
}

func updateDB(invitation Invitation) {
	db := database.GetDB()
	db.Model(&invitation).Where("id = ?", invitation.ID).Update("message", invitation.Message).Update("accepted", invitation.Accepted)
}

func deleteDB(ID int) {
	db := database.GetDB()
	var invitation Invitation
	invitation.ID = ID
	db.Delete(&invitation)
}

func deleteByUserDB(userID int) {
	db := database.GetDB()
	db.Delete(&Invitation{}, "invitee_id = ?", userID)
}

func deleteByAppointmentDB(appointmentID int) {
	db := database.GetDB()
	db.Delete(&Invitation{}, "appointment_id = ?", appointmentID)
}
