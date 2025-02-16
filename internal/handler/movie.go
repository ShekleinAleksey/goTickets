package handler

import (
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie.ID = len(movies) + 1
	movies = append(movies, movie)

	c.JSON(http.StatusCreated, movie)
}

func (h *Handler) GetMovies(c *gin.Context) {
	movies, err := h.service.MovieService.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}
