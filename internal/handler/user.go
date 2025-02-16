package handler

import (
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Создание пользователя
	id, err := h.service.UserService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.ID = id

	c.JSON(http.StatusCreated, user)
}

// Пополнить баланс
func (h *Handler) ReplenishBalance(c *gin.Context) {

}
