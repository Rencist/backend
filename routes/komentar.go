package routes

import (
	"tamiyochi-backend/controller"
	"tamiyochi-backend/middleware"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

func KomentarRoutes(router *gin.Engine, KomentarController controller.KomentarController, jwtService service.JWTService) {
	komentarRoutes := router.Group("/api/komentar")
	{
		komentarRoutes.POST("", middleware.Authenticate(jwtService, false), KomentarController.CreateKomentar)
		komentarRoutes.GET("/:id", KomentarController.FindKomentarBySeriID)
	}
}