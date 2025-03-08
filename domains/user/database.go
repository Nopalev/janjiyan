package user

import (
	"github.com/Nopalev/janjiyan/utilities/database"
	"gorm.io/gorm/clause"
)

func createDB(user User) User {
	db := database.GetDB()

	if db != nil {
		db.Create(&user)
	}
	return user
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

func updateDB(username string, user User) User {
	db := database.GetDB()
	db.Model(&User{}).Clauses(clause.Returning{}).Where("username = ?", username).Updates(&user)
	return user
}

func deleteDB(username string) {
	db := database.GetDB()
	var user User
	db.First(&user, User{Username: username})
	db.Delete(&user)
}
