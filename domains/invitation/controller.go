package invitation

import (
	"errors"
	"log"
	"slices"

	"github.com/Nopalev/janjiyan/domains/appointment"
	"github.com/Nopalev/janjiyan/domains/user"
)

func Create(invitation Invitation, issuer string) (Invitation, error) {
	checkedAppointment, err := appointment.Read(invitation.AppointmentID, issuer)
	log.Println(checkedAppointment, issuer)
	if err != nil {
		return invitation, err
	}

	if user.IDbyUsername(issuer) == invitation.InviteeID {
		return invitation, errors.New("you can't invite yourself")
	}

	createDB(&invitation)
	invitation.Appointment = checkedAppointment
	assocUser(&invitation)
	invitation.UserWithoutPassword = user.RemovePassword(invitation.User)
	return invitation, nil
}

func Read(ID int, issuer string) (Invitation, error) {
	invitation := readDB(ID)
	assocAppointment(&invitation)
	assocUser(&invitation)

	if !(invitation.Appointment.User.Username == issuer ||
		invitation.User.Username == issuer) {
		return invitation, errors.New("you are not involved in this appointment")
	}
	invitation.UserWithoutPassword = user.RemovePassword(invitation.User)
	return invitation, nil
}

func ReadByAppointment(appointment_ID int, issuer string) (Member, error) {
	invitations := readByAppointmentDB(appointment_ID)
	members := getMember(invitations)

	if !(members.Creator == issuer ||
		slices.Index(members.Accepted, issuer) != -1 ||
		slices.Index(members.Invited, issuer) != -1) {
		return members, errors.New("you are no involved in this appointment")
	}

	return members, nil
}

func ReadByCreator(issuer string) []Invitation {
	var appointmentsID []int
	appointments := appointment.ReadCreated(issuer)
	for _, val := range appointments {
		appointmentsID = append(appointmentsID, val.ID)
	}
	invitations := readByAppointmentsDB(appointmentsID)
	for idx, _ := range invitations {
		assocAppointment(&invitations[idx])
		assocUser(&invitations[idx])
		invitations[idx].UserWithoutPassword = user.RemovePassword(invitations[idx].User)
	}
	return invitations
}

func InvitedAppointments(issuer string) []InvitationWithoutUser {
	var invitationsWithoutUser []InvitationWithoutUser
	invitations := readByUserDB(user.IDbyUsername(issuer))
	for idx, _ := range invitations {
		assocAppointment(&invitations[idx])
		invitationsWithoutUser = append(invitationsWithoutUser, removeUserInformation(invitations[idx]))
	}
	return invitationsWithoutUser
}

func Update(invitation Invitation, issuer string, accept bool) (InvitationWithoutUser, error) {
	check := readDB(invitation.ID)
	if accept {
		assocUser(&check)
		if check.User.Username != issuer {
			return removeUserInformation(invitation), errors.New("you are not the invitee of this invitation")
		}
		check.Accepted = true
	} else {
		assocAppointment(&check)
		if check.Appointment.User.Username != issuer {
			return removeUserInformation(invitation), errors.New("you are not the creator of this invitation")
		}
		check.Message = invitation.Message
	}

	updateDB(check)
	updatedInvitation := readDB(invitation.ID)
	assocAppointment(&updatedInvitation)
	return removeUserInformation(updatedInvitation), nil
}

func Delete(ID int, issuer string) error {
	check := readDB(ID)
	assocAppointment(&check)
	if check.Appointment.User.Username != issuer {
		return errors.New("you are not the creator of this invitation")
	}
	deleteDB(ID)
	return nil
}

func DeleteByUser(userID int) {
	deleteByUserDB(userID)
}

func DeleteByAppointment(appointmentID int) {
	deleteByAppointmentDB(appointmentID)
}
