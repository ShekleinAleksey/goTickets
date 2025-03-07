package handler

import (
	"net/http"
	"strconv"

	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/gin-gonic/gin"
)

type MovieService interface {
	GetAllMovies() ([]entity.Movie, error)
	GetMovieByID(id int) (entity.Movie, error)
}

type MovieHandler struct {
	MovieService MovieService
}

func NewMovieHandler(movieService MovieService) *MovieHandler {
	return &MovieHandler{MovieService: movieService}
}

// @Summary Get Movies
// @Tags movie
// @Description get all movies
// @ID get-movies
// @Accept json
// @Produce json
// @Success 200 {object} entity.Movie
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /movies [get]
func (h *MovieHandler) GetMovies(c *gin.Context) {
	movies, err := h.MovieService.GetAllMovies()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, movies)
}

// @Summary Get Movie By ID
// @Tags movie
// @Description get movie by id
// @ID get-movie-by-id
// @Accept json
// @Produce json
// @Param id path int true "movie id"
// @Success 200 {object} entity.Movie
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /movies/{id} [get]
func (h *MovieHandler) GetMovieByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	movie, err := h.MovieService.GetMovieByID(idInt)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, movie)
}
