package routes

import (
	"tamiyochi-backend/controller"
	"tamiyochi-backend/middleware"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

func MangaRoutes(router *gin.Engine, MangaController controller.MangaController, jwtService service.JWTService) {
	mangaRoutes := router.Group("/api/manga")
	{
		mangaRoutes.POST("", MangaController.CreateManga)
		mangaRoutes.GET("", middleware.Authenticate(jwtService, false), MangaController.GetAllManga)
		mangaRoutes.DELETE("/", middleware.Authenticate(jwtService, false), MangaController.DeleteManga)
		mangaRoutes.PUT("/", middleware.Authenticate(jwtService, false), MangaController.UpdateManga)
	}
}