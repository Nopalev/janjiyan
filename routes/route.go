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

	userRoutes := r.Group("/user").Use(authMiddleware)
	{
		userRoutes.PUT("/update", userUpdate)
		userRoutes.DELETE("/delete", userDelete)
	}

	appointmentRoutes := r.Group("/appointment").Use(authMiddleware)
	{
		appointmentRoutes.POST("/create", appointmentCreate)
		appointmentRoutes.GET("/:id", appointmentGet)
		appointmentRoutes.PUT("/update", appointmentUpdate)
		appointmentRoutes.DELETE("/delete", appointmentDelete)
	}

	appointmentsRoutes := r.Group("/appointments").Use(authMiddleware)
	{
		appointmentsRoutes.GET("/created", appointmentCreated)
	}

	return r
}
