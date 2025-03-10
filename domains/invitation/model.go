package invitation

import (
	"github.com/Nopalev/janjiyan/domains/appointment"
	"github.com/Nopalev/janjiyan/domains/user"
)

type Invitation struct {
	ID                  int                      `json:"ID" gorm:"primaryKey"`
	Message             string                   `json:"message"`
	AppointmentID       int                      `json:"appointment_id" gorm:"not null"`
	Appointment         appointment.Appointment  `json:"appointment"`
	InviteeID           int                      `json:"invitee_id" gorm:"not null"`
	Accepted            bool                     `json:"accepted"`
	User                user.User                `json:"-" gorm:"foreignKey:InviteeID"`
	UserWithoutPassword user.UserWithoutPassword `json:"invited_user" gorm:"-"`
}

type InvitationWithoutUser struct {
	ID            int                     `json:"ID"`
	Message       string                  `json:"message"`
	AppointmentID int                     `json:"appointment_id"`
	Appointment   appointment.Appointment `json:"appointment"`
	InviteeID     int                     `json:"invitee_id"`
	Accepted      bool                    `json:"accepted"`
}

type Member struct {
	Appointment appointment.Appointment `json:"appointment"`
	Creator     string                  `json:"creator"`
	Accepted    []string                `json:"accepted"`
	Invited     []string                `json:"invited"`
}

func getMember(invitations []Invitation) Member {
	var member Member
	assocAppointment(&invitations[0])
	member.Appointment = invitations[0].Appointment
	member.Creator = invitations[0].Appointment.User.Username
	for _, val := range invitations {
		assocUser(&val)
		if val.Accepted {
			member.Accepted = append(member.Accepted, val.User.Username)
		} else {
			member.Invited = append(member.Invited, val.User.Username)
		}
	}
	return member
}

func removeUserInformation(invitation Invitation) InvitationWithoutUser {
	return InvitationWithoutUser{
		ID:            invitation.ID,
		Message:       invitation.Message,
		AppointmentID: invitation.AppointmentID,
		Appointment:   invitation.Appointment,
		InviteeID:     invitation.InviteeID,
		Accepted:      invitation.Accepted,
	}
}
