package handler

import (
	_ "github.com/ShekleinAleksey/goTickets/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Deps struct {
	MovieService          MovieService
	MovieScreeningService MovieScreeningService
	TicketService         TicketService
	UserService           UserService
}

type Handler struct {
	MovieHandler          *MovieHandler
	MovieScreeningHandler *MovieScreeningHandler
	TicketHandler         *TicketHandler
	UserHandler           *UserHandler
}

func NewHandler(deps Deps) *Handler {
	return &Handler{
		MovieHandler:          NewMovieHandler(deps.MovieService),
		MovieScreeningHandler: NewMovieSceeningHandler(deps.MovieScreeningService),
		TicketHandler:         NewTicketHandler(deps.TicketService),
		UserHandler:           NewUserHandler(deps.UserService),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		movies := api.Group("/movies")
		{
			movies.GET("/", h.MovieHandler.GetMovies)
			movies.GET("/:id", h.MovieHandler.GetMovieByID)
		}

		users := api.Group("/users")
		{
			users.POST("/", h.UserHandler.CreateUser)
			users.POST("/:id/replenish-balance", h.UserHandler.ReplenishBalance)
		}

		screenings := api.Group("/screenings")
		{
			screenings.GET("/", h.MovieScreeningHandler.GetMovieScreenings)
			screenings.GET("/:id", h.MovieScreeningHandler.GetMovieScreening)
		}

		tickets := api.Group("/tickets")
		{
			tickets.GET("/:id", h.TicketHandler.GetTicket)
			tickets.POST("/buy", h.TicketHandler.BuyTicket)
		}
	}

	return router
}
