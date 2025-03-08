package user

type User struct {
	ID       int    `json:"ID" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Username string `json:"username" gorm:"not null;unique"`
	Timezone string `json:"timezone"`
	Password string `json:"password" gorm:"not null"`
}
