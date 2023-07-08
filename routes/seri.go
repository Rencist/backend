package routes

import (
	"tamiyochi-backend/controller"
	"tamiyochi-backend/middleware"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

func SeriRoutes(router *gin.Engine, SeriController controller.SeriController, jwtService service.JWTService) {
	seriRoutes := router.Group("/api/seri")
	{
		seriRoutes.POST("", middleware.Authenticate(jwtService, true), SeriController.CreateSeri)
		seriRoutes.GET("", SeriController.GetAllSeri)
		seriRoutes.GET("/:id", SeriController.FindSeriByID)
		seriRoutes.PUT("/:id", middleware.Authenticate(jwtService, true), SeriController.UpdateSeri)
		seriRoutes.DELETE("/", middleware.Authenticate(jwtService, false), SeriController.DeleteSeri)
		seriRoutes.PUT("/", middleware.Authenticate(jwtService, false), SeriController.UpdateSeri)
		seriRoutes.POST("/rating", middleware.Authenticate(jwtService, false), SeriController.UpsertRating)
		seriRoutes.GET("/rating", middleware.Authenticate(jwtService, false), SeriController.GetRating)
		seriRoutes.GET("/rating/:seri_id", middleware.Authenticate(jwtService, false), SeriController.GetRatingSeri)
	}
}