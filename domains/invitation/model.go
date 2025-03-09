package invitation

import (
	"github.com/Nopalev/janjiyan/domains/appointment"
	"github.com/Nopalev/janjiyan/domains/user"
)

type Invitation struct {
	ID            int                     `json:"ID" gorm:"primaryKey"`
	Message       string                  `json:"message"`
	AppointmentID int                     `json:"appointment_id" gorm:"not null"`
	Appointment   appointment.Appointment `json:"-"`
	InviteeID     int                     `json:"inviteeID" gorm:"not null"`
	Accepted      bool                    `json:"accepted"`
	User          user.User               `json:"-" gorm:"foreignKey:InviteeID"`
}
