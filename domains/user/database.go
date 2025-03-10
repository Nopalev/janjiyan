package user

import (
	"github.com/Nopalev/janjiyan/utilities/database"
)

func createDB(user *User) error {
	db := database.GetDB()
	res := db.Create(&user)

	return res.Error
}

func readDB(username string) User {
	db := database.GetDB()
	var user User
	db.First(&user, User{Username: username})
	return user
}

func userExistDB(username string) bool {
	db := database.GetDB()
	var user User
	res := db.Limit(1).Find(&user, User{Username: username})
	return res.RowsAffected == 1
}

func updateDB(username string, user User) error {
	db := database.GetDB()
	res := db.Model(&user).Where("username = ?", username).Updates(user)
	return res.Error
}

func deleteDB(username string) {
	db := database.GetDB()
	var user User
	db.First(&user, User{Username: username})
	db.Delete(&user)
}
