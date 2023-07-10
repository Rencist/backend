package routes

import (
	"tamiyochi-backend/controller"
	"tamiyochi-backend/middleware"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(router *gin.Engine, MovieController controller.MovieController, jwtService service.JWTService) {
	movieRoutes := router.Group("/api/movie")
	{
		movieRoutes.GET("", MovieController.GetAllMovie)
		movieRoutes.GET("/:id", MovieController.GetMovieByID)
		movieRoutes.POST("/transaction", middleware.Authenticate(jwtService, false), MovieController.CreateTransaction)
		movieRoutes.GET("/available_seat/:id", MovieController.GetAvailableSeat)
	}
}