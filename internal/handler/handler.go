package handler

import (
	_ "github.com/ShekleinAleksey/goTickets/docs"
	"github.com/ShekleinAleksey/goTickets/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		movies := api.Group("/movies")
		{
			movies.POST("/", h.CreateMovie)
			movies.GET("/", h.GetMovies)
			movies.GET("/:id", h.GetMovieByID)
		}

		users := api.Group("/users")
		{
			users.POST("/", h.CreateUser)
			users.POST("/:id/replenish-balance", h.ReplenishBalance)
		}

		screenings := api.Group("/screenings")
		{
			screenings.POST("/movie", h.CreateMovieScreening)
			screenings.GET("/", h.GetMovieScreenings)
			screenings.GET("/:id", h.GetMovieScreening)
		}

		tickets := api.Group("/tickets")
		{
			tickets.GET("/:id", h.GetTicket)
			tickets.POST("/buy", h.BuyTicket)
		}
	}

	return router
}
