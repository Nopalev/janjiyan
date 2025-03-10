package user

type User struct {
	ID       int    `json:"ID" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Username string `json:"username" gorm:"not null;unique;uniqueIndex"`
	Timezone int    `json:"timezone"`
	Password string `json:"password" gorm:"not null"`
}

type UserWithoutPassword struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Timezone int    `json:"timezone"`
}

func RemovePassword(user User) UserWithoutPassword {
	return UserWithoutPassword{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Timezone: user.Timezone,
	}
}
