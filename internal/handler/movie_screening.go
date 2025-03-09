package handler

import (
	"net/http"
	"strconv"

	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/gin-gonic/gin"
)

type MovieScreeningService interface {
	GetAllMovieScreenings() ([]entity.MovieScreening, error)
	GetScreening(id int) (entity.MovieScreening, error)
	UpdateScreening(screening *entity.MovieScreening) error
}

type MovieScreeningHandler struct {
	MovieScreeningService MovieScreeningService
}

func NewMovieSceeningHandler(movieScreeningService MovieScreeningService) *MovieScreeningHandler {
	return &MovieScreeningHandler{MovieScreeningService: movieScreeningService}
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
func (h *MovieScreeningHandler) GetMovieScreenings(c *gin.Context) {
	screenings, err := h.MovieScreeningService.GetAllMovieScreenings()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	for i, screening := range screenings {
		movie, err := h.service.MovieService.GetMovieByID(screening.MovieID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		screenings[i].Movie = movie
	}
	c.JSON(http.StatusOK, screenings)
}

// @Summary Get Movie Screening
// @Tags movie_screening
// @Description get movie screening
// @ID get-movie-screening
// @Produce json
// @Param id path int true "Movie Screening ID"
// @Success 200 {object} entity.MovieScreening
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /screenings/{id} [get]
func (h *MovieScreeningHandler) GetMovieScreening(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid screening ID")
		return
	}
	screening, err := h.MovieScreeningService.GetScreening(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, screening)
}

// @Summary Filter Movie Screenings
// @Tags movie_screening
// @Description get movie screenings filtered by movie ID and date
// @ID filter-movie-screenings
// @Produce json
// @Param movie_id query int false "Movie ID"
// @Param date query string false "Date (YYYY-MM-DD)"
// @Success 200 {object} []entity.MovieScreening
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /screenings/filter [get]
func (h *MovieScreeningHandler) FilterMovieScreenings(c *gin.Context) {

	movieIDStr := c.Query("movie_id")
	movieID, err := strconv.Atoi(movieIDStr)
	dateStr := c.Query("date")

	screenings, err := h.MovieScreeningService.FilterScreenings(movieID, dateStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, screenings)
}
