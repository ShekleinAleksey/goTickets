package handler

import (
	"net/http"

	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/gin-gonic/gin"
)

var movieScreenings []entity.MovieScreening

// @Summary Create Movie Screening
// @Tags movie_screening
// @Description create movie screening
// @ID create-movie-screening
// @Accept json
// @Produce json
// @Param input body entity.MovieScreening true "create movie screening"
// @Success 200 {object} entity.MovieScreening
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /screenings [post]
func (h *Handler) CreateMovieScreening(c *gin.Context) {
	var screening entity.MovieScreening
	if err := c.BindJSON(&screening); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if err := json.NewDecoder(r.Body).Decode(&screening); err != nil {
	// 	http.Error(w, "Invalid request body", http.StatusBadRequest)
	// 	return
	// }

	screening.ID = len(movieScreenings) + 1
	movieScreenings = append(movieScreenings, screening)

	c.JSON(http.StatusCreated, screening)
}

// @Summary Get Movie Screenings
// @Tags movie_screening
// @Description get movie screenings
// @ID get-movie-screenings
// @Produce json
// @Success 200 {object} []entity.MovieScreening
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /screenings [get]
func (h *Handler) GetMovieScreenings(c *gin.Context) {
	screenings, err := h.service.MovieScreeningService.GetAllMovieScreenings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, screenings)
}
