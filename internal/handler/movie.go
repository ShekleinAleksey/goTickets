package handler

import (
	"net/http"
	"strconv"

	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/gin-gonic/gin"
)

var movies []entity.Movie

// @Summary Create Movie
// @Tags movie
// @Description create movie
// @ID create-movie
// @Accept json
// @Produce json
// @Param input body entity.Movie true "create movie"
// @Success 200 {object} entity.Movie
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /movies [post]
func (h *Handler) CreateMovie(c *gin.Context) {
	var movie entity.Movie
	if err := c.BindJSON(&movie); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	movie.ID = len(movies) + 1
	movies = append(movies, movie)

	c.JSON(http.StatusCreated, movie)
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
func (h *Handler) GetMovies(c *gin.Context) {
	movies, err := h.service.MovieService.GetAllMovies()
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
func (h *Handler) GetMovieByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	movie, err := h.service.MovieService.GetMovieByID(idInt)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, movie)
}
