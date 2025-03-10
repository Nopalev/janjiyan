package appointment

import (
	"errors"
	"log"

	"github.com/Nopalev/janjiyan/domains/user"
)

func Create(appointment Appointment, issuer string) Appointment {
	appointment.CreatorID = user.IDbyUsername(issuer)
	createDB(&appointment)
	return appointment
}

func Read(ID int, issuer string) (Appointment, error) {
	appointment := readDB(ID)
	log.Println(appointment)
	if issuer != appointment.User.Username {
		return appointment, errors.New("you are not the creator of this appointment")
	}
	return appointment, nil
}

func ReadCreated(username string) []Appointment {
	return readCreatedDB(user.IDbyUsername(username))
}

func Update(issuer string, appointment Appointment) (Appointment, error) {
	check := readDB(appointment.ID)
	if issuer != check.User.Username {
		return appointment, errors.New("you are not the creator of this appointment")
	}
	updateDB(appointment)
	appointment = readDB(appointment.ID)
	return appointment, nil
}

func Delete(issuer string, invitationDeletion func(int), appointment Appointment) error {
	check := readDB(appointment.ID)
	if issuer != check.User.Username {
		return errors.New("you are not the creator of this appointment")
	}
	invitationDeletion(appointment.ID)
	deleteDB(appointment.ID)
	return nil
}

func DeleteByUser(userID int) {
	deleteByUserDB(userID)
}
