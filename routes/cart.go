package routes

import (
	"tamiyochi-backend/controller"
	"tamiyochi-backend/middleware"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

func CartRoutes(router *gin.Engine, CartController controller.CartController, jwtService service.JWTService) {
	cartRoutes := router.Group("/api/cart")
	{
		cartRoutes.POST("", middleware.Authenticate(jwtService, false), CartController.CreateCart)
		cartRoutes.GET("", middleware.Authenticate(jwtService, false), CartController.FindCartByUserID)
		cartRoutes.DELETE("/:manga_id", middleware.Authenticate(jwtService, false), CartController.DeleteCart)
		cartRoutes.DELETE("/manga/:manga_id", middleware.Authenticate(jwtService, false), CartController.DeleteAllByMangaIDCart)
	}
}