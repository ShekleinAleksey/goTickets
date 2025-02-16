package handler

import (
	"github.com/ShekleinAleksey/goTickets/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		movies := api.Group("/movies")
		{
			movies.POST("/", h.CreateMovie)
			movies.GET("/", h.GetMovies)
		}

		users := api.Group("/users")
		{
			users.POST("/", h.CreateUser)
		}

		screenings := api.Group("/screenings")
		{
			screenings.POST("/movie", h.CreateMovieScreening)
		}

		tickets := api.Group("/tickets")
		{
			tickets.GET("/get", h.GetTicket)
			tickets.POST("/buy", h.BuyTicket)
			tickets.GET("/user/:id", h.GetUserTickets)
		}
	}

	return router
}
