package routes

import (
	"tamiyochi-backend/controller"

	"github.com/gin-gonic/gin"
)

func ProvinsiRoutes(router *gin.Engine, ProvinsiController controller.ProvinsiController) {
	provinsiRoutes := router.Group("/api/provinsi")
	{
		provinsiRoutes.GET("", ProvinsiController.GetAllProvinsi)
	}
}