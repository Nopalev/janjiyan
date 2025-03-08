package user

import (
	"errors"
	"log"

	"github.com/Nopalev/janjiyan/utilities/auth"
)

func Register(user User) (User, string) {
	var err error
	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
	}
	user = createDB(user)
	token, err := auth.CreateToken(user.Username)
	if err != nil {
		log.Println(err)
	}
	return user, token
}

func Login(attemptedUser User) (string, error) {
	var token string
	var err error
	user := readDB(attemptedUser.Username)
	if auth.CheckPasswordHash(attemptedUser.Password, user.Password) {
		token, err = auth.CreateToken(attemptedUser.Username)
		if err != nil {
			return "", err
		}
	} else {
		err = errors.New("invalid credentials")
	}
	return token, err
}

func CheckIfUserExist(username string) bool {
	return userExistDB(username)
}

func Update(username string, user User) (User, string) {
	if user.Password != "" {
		var err error
		user.Password, err = auth.HashPassword(user.Password)

		if err != nil {
			log.Println(err)
		}
	}
	updateDB(username, user)
	user = readDB(user.Username)
	token, err := auth.CreateToken(user.Username)
	if err != nil {
		return user, ""
	}
	return user, token
}

func Delete(username string) {
	deleteDB(username)
}
