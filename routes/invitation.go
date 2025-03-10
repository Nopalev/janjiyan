package routes

import (
	"net/http"
	"strconv"

	"github.com/Nopalev/janjiyan/domains/invitation"
	"github.com/gin-gonic/gin"
)

func appointmentMembers(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	issuer := ctx.MustGet("issuer").(string)
	members, err := invitation.ReadByAppointment(ID, issuer)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &members)
}

func invitedAppointments(ctx *gin.Context) {
	issuer := ctx.MustGet("issuer").(string)
	invited := invitation.InvitedAppointments(issuer)
	ctx.JSON(http.StatusOK, &invited)
}

func getInvitation(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	issuer := ctx.MustGet("issuer").(string)
	invitation, err := invitation.Read(ID, issuer)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &invitation)
}

func createdInvitations(ctx *gin.Context) {
	issuer := ctx.MustGet("issuer").(string)
	invitations := invitation.ReadByCreator(issuer)
	ctx.JSON(http.StatusOK, &invitations)
}

func createInvitation(ctx *gin.Context) {
	issuer := ctx.MustGet("issuer").(string)
	var newInvitation invitation.Invitation
	err := ctx.BindJSON(&newInvitation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newInvitation, err = invitation.Create(newInvitation, issuer)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &newInvitation)
}

func acceptInvitation(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	issuer := ctx.MustGet("issuer").(string)
	acceptedInvitation, err := invitation.Update(
		invitation.Invitation{
			ID: ID,
		},
		issuer,
		true,
	)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &acceptedInvitation)
}

func updateInvitation(ctx *gin.Context) {
	issuer := ctx.MustGet("issuer").(string)
	var invitationInformation invitation.Invitation
	err := ctx.BindJSON(&invitationInformation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedInvitation, err := invitation.Update(invitationInformation, issuer, false)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &updatedInvitation)
}

func deleteInvitation(ctx *gin.Context) {
	var invitationToBeDeleted invitation.Invitation
	err := ctx.BindJSON(&invitationToBeDeleted)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	issuer := ctx.MustGet("issuer").(string)
	err = invitation.Delete(invitationToBeDeleted.ID, issuer)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "invitation deleted",
	})
}
