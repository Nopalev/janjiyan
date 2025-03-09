package appointment

import (
	"errors"

	"github.com/Nopalev/janjiyan/domains/user"
)

func Create(appointment Appointment, issuer string) Appointment {
	appointment.CreatorID = user.IDbyUsername(issuer)
	return createDB(appointment)
}

func Read(ID int, issuer string) (Appointment, error) {
	appointment := readDB(ID)
	if user.IDbyUsername(issuer) != appointment.CreatorID {
		return appointment, errors.New("you are not the creator of this appointment")
	}
	return appointment, nil
}

func ReadCreated(username string) []Appointment {
	return readCreatedDB(user.IDbyUsername(username))
}

func Update(issuer string, appointment Appointment) (Appointment, error) {
	check := readDB(appointment.ID)
	if user.IDbyUsername(issuer) != check.CreatorID {
		return appointment, errors.New("you are not the creator of this appointment")
	}
	updateDB(appointment)
	appointment = readDB(appointment.ID)
	return appointment, nil
}

func Delete(issuer string, appointment Appointment) error {
	check := readDB(appointment.ID)
	if user.IDbyUsername(issuer) != check.CreatorID {
		return errors.New("you are not the creator of this appointment")
	}
	deleteDB(appointment.ID)
	return nil
}
