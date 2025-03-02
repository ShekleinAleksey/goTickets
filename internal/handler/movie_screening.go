package handler

import (
	"net/http"
	"strconv"
	"time"

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
		newErrorResponse(c, http.StatusBadRequest, err.Error())
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
func (h *Handler) GetMovieScreening(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid screening ID")
		return
	}
	screening, err := h.service.MovieScreeningService.GetScreening(id)
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
func (h *Handler) FilterMovieScreenings(c *gin.Context) {
	// Получаем параметры фильтрации
	movieIDStr := c.Query("movie_id")
	dateStr := c.Query("date")

	// Получаем все сеансы
	screenings, err := h.service.MovieScreeningService.GetAllMovieScreenings()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Фильтруем по ID фильма, если указан
	if movieIDStr != "" {
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Invalid movie ID")
			return
		}

		filteredScreenings := make([]entity.MovieScreening, 0)
		for _, screening := range screenings {
			if screening.MovieID == movieID {
				filteredScreenings = append(filteredScreenings, screening)
			}
		}
		screenings = filteredScreenings
	}

	// Фильтруем по дате, если указана
	if dateStr != "" {
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD")
			return
		}

		filteredScreenings := make([]entity.MovieScreening, 0)
		for _, screening := range screenings {
			if screening.ShowTime.Format("2006-01-02") == date.Format("2006-01-02") {
				filteredScreenings = append(filteredScreenings, screening)
			}
		}
		screenings = filteredScreenings
	}

	// Загружаем информацию о фильмах для каждого сеанса
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
