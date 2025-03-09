package routes

import (
	"net/http"
	"strconv"

	"github.com/Nopalev/janjiyan/domains/appointment"
	"github.com/gin-gonic/gin"
)

func appointmentCreate(ctx *gin.Context) {
	var newAppointment appointment.Appointment
	err := ctx.BindJSON(&newAppointment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	issuer := ctx.MustGet("issuer").(string)

	newAppointment = appointment.Create(newAppointment, issuer)
	ctx.JSON(http.StatusOK, &newAppointment)
}

func appointmentGet(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	issuer := ctx.MustGet("issuer").(string)
	appointment, err := appointment.Read(ID, issuer)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &appointment)
}

func appointmentCreated(ctx *gin.Context) {
	issuer := ctx.MustGet("issuer").(string)
	appointments := appointment.ReadCreated(issuer)
	ctx.JSON(http.StatusOK, &appointments)
}

func appointmentUpdate(ctx *gin.Context) {
	var updatedAppointment appointment.Appointment
	err := ctx.BindJSON(&updatedAppointment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	issuer := ctx.MustGet("issuer").(string)
	updatedAppointment, err = appointment.Update(issuer, updatedAppointment)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &updatedAppointment)
}

func appointmentDelete(ctx *gin.Context) {
	var appointmentToBeDeleted appointment.Appointment
	err := ctx.BindJSON(&appointmentToBeDeleted)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	issuer := ctx.MustGet("issuer").(string)
	err = appointment.Delete(issuer, appointmentToBeDeleted)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "appointment deleted",
	})
}
