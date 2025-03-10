package user

import (
	"errors"
	"log"

	"github.com/Nopalev/janjiyan/utilities/auth"
)

func Register(user User) (UserWithoutPassword, string, error) {
	var err error
	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
	}
	err = createDB(&user)
	if err != nil {
		return RemovePassword(user), "", errors.New("username cannot be used")
	}
	token, err := auth.CreateToken(user.Username)
	if err != nil {
		log.Println(err)
	}
	user.Password = ""
	return RemovePassword(user), token, nil
}

func IDbyUsername(username string) int {
	return readDB(username).ID
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

func Update(username string, user User) (UserWithoutPassword, string, error) {
	if user.Password != "" {
		var err error
		user.Password, err = auth.HashPassword(user.Password)

		if err != nil {
			log.Println(err)
		}
	}

	err := updateDB(username, user)
	if err != nil {
		return RemovePassword(user), "", errors.New("username cannot be used")
	}

	user = readDB(username)
	token, err := auth.CreateToken(user.Username)
	if err != nil {
		return RemovePassword(user), "", err
	}
	user.Password = ""
	return RemovePassword(user), token, nil
}

func Delete(username string, invitationDeletion func(int), appointmentDeletion func(int)) {
	invitationDeletion(IDbyUsername(username))
	appointmentDeletion(IDbyUsername(username))
	deleteDB(username)
}
