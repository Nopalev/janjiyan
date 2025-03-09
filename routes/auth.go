package routes

import (
	"net/http"

	"github.com/Nopalev/janjiyan/domains/user"
	"github.com/gin-gonic/gin"
)

func register(ctx *gin.Context) {
	var newUser user.User
	var token string
	err := ctx.BindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newUser, token, err = user.Register(newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"user":  &newUser,
		"token": token,
	})
}

func login(ctx *gin.Context) {
	var attemptedUser user.User
	err := ctx.BindJSON(&attemptedUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := user.Login(attemptedUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
