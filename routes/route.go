package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", register)
		authRoutes.GET("/login", login)
	}

	userRoutes := r.Group("/user").Use(middleware)
	{
		userRoutes.PUT("/update", userUpdate)
		userRoutes.DELETE("/delete", userDelete)
	}

	return r
}
