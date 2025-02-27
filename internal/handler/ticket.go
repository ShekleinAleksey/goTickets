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
// @Param id path int true "Ticket ID"
// @Success 200 {object} entity.Ticket
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /tickets/{id} [get]
func (h *Handler) GetTicket(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ticket ID")
		return
	}

	ticket, err := h.service.TicketService.GetTicket(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ticket)
}

// @Summary Buy Ticket
// @Tags ticket
// @Description buy ticket
// @ID buy-ticket
// @Accept json
// @Produce json
// @Param input body entity.Ticket true "buy ticket"
// @Success 200 {object} entity.Ticket
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /tickets/buy [post]
func (h *Handler) BuyTicket(c *gin.Context) {
	idStr := c.Query("screening_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid screening ID")
		return
	}

	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	seatStr := c.Query("seat")
	seat, err := strconv.Atoi(seatStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid seat")
		return
	}

	screening, err := h.service.MovieScreeningService.GetScreening(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.TicketService.BuyTicket(screening, userID, seat)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.service.UserService.GetUserByID(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.Balance -= screening.TicketPrice
	err = h.service.UserService.UpdateUser(&user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket bought successfully"})
}
