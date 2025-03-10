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
		userRoutes.PUT("/update", updateUser)
		userRoutes.DELETE("/delete", deleteUser)
	}

	appointmentRoutes := r.Group("/appointment").Use(authMiddleware)
	{
		appointmentRoutes.POST("/create", createAppointment)
		appointmentRoutes.GET("/:id", getAppointment)
		appointmentRoutes.GET("/:id/members", appointmentMembers)
		appointmentRoutes.PUT("/update", updateAppointment)
		appointmentRoutes.DELETE("/delete", deleteAppointment)
	}

	appointmentsRoutes := r.Group("/appointments").Use(authMiddleware)
	{
		appointmentsRoutes.GET("/created", createdAppointments)
		appointmentsRoutes.GET("/invited", invitedAppointments)
	}

	invitationRoutes := r.Group("/invitation").Use(authMiddleware)
	{
		invitationRoutes.POST("/create", createInvitation)
		invitationRoutes.GET("/:id", getInvitation)
		invitationRoutes.POST("/:id/accept", acceptInvitation)
		invitationRoutes.PUT("/update", updateInvitation)
		invitationRoutes.DELETE("/delete", deleteInvitation)
	}

	invitationsRoutes := r.Group("/invitations").Use(authMiddleware)
	{
		invitationsRoutes.GET("/created", createdInvitations)
	}

	return r
}
