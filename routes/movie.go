package routes

import (
	"tamiyochi-backend/controller"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(router *gin.Engine, MovieController controller.MovieController) {
	movieRoutes := router.Group("/api/movie")
	{
		movieRoutes.GET("", MovieController.GetAllMovie)
		movieRoutes.GET("/:id", MovieController.GetMovieByID)
	}
}