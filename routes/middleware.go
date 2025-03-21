package routes

import (
	"net/http"

	"github.com/Nopalev/janjiyan/domains/user"
	"github.com/Nopalev/janjiyan/utilities/auth"
	"github.com/gin-gonic/gin"
)

func authMiddleware(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "no token is being found",
		})
		return
	}

	tokenString = tokenString[len("Bearer "):]

	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	issuer, err := token.Claims.GetIssuer()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	check := user.CheckIfUserExist(issuer)
	if !check {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "user unregistered",
		})
		return
	}

	ctx.Set("issuer", issuer)

	ctx.Next()
}
