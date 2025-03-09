package appointment

import (
	"time"

	"github.com/Nopalev/janjiyan/domains/user"
)

type Appointment struct {
	ID        int       `json:"ID" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Start     time.Time `json:"start" gorm:"not null"`
	End       time.Time `json:"end" gorm:"not null"`
	CreatorID int       `json:"creatorID" gorm:"not null"`
	User      user.User `json:"-" gorm:"foreignKey:CreatorID"`
}
