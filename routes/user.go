package routes

import (
	"net/http"

	"github.com/Nopalev/janjiyan/domains/user"
	"github.com/gin-gonic/gin"
)

func userUpdate(ctx *gin.Context) {
	var updatedUser user.User
	err := ctx.BindJSON(&updatedUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	issuer := ctx.MustGet("issuer").(string)
	user, token, err := user.Update(issuer, updatedUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"user":  &user,
		"token": token,
	})
}

func userDelete(ctx *gin.Context) {
	issuer := ctx.MustGet("issuer").(string)
	user.Delete(issuer)
}
