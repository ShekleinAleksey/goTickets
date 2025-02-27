package handler

import (
	"net/http"
	"strconv"

	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/gin-gonic/gin"
)

// @Summary Create User
// @Tags user
// @Description create user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body entity.User true "create user"
// @Success 200 {object} entity.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Создание пользователя
	id, err := h.service.UserService.CreateUser(&user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.ID = id

	c.JSON(http.StatusCreated, user)
}

// @Summary Replenish Balance
// @Tags user
// @Description replenish balance
// @ID replenish-balance
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body int true "Balance"
// @Success 200 {object} entity.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id}/replenish-balance [post]
func (h *Handler) ReplenishBalance(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input int
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.UserService.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.Balance += float64(input)
	err = h.service.UserService.UpdateUser(&user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Balance replenished successfully"})
}
