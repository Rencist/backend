package routes

import (
	"tamiyochi-backend/controller"

	"github.com/gin-gonic/gin"
)

func KabupatenRoutes(router *gin.Engine, KabupatenController controller.KabupatenController) {
	kabupatenRoutes := router.Group("/api/kabupaten")
	{
		kabupatenRoutes.GET("", KabupatenController.GetAllKabupaten)
		kabupatenRoutes.GET("/:id", KabupatenController.FindKabupatenByProvinsiID)
	}
}