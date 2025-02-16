package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get Ticket
// @Tags ticket
// @Description get ticket
// @ID get-ticket
// @Accept json
// @Produce json
// @Param input body entity.Ticket true "get ticket"
// @Success 200 {object} entity.Ticket
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /tickets/get [get]
func (h *Handler) GetTicket(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	ticket, err := h.service.TicketService.GetTicket(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func (h *Handler) BuyTicket(c *gin.Context) {

}

func (h *Handler) GetUserTickets(c *gin.Context) {

}
