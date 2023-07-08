package controller

import (
	"net/http"
	"tamiyochi-backend/common"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

type ProvinsiController interface {
	GetAllProvinsi(ctx *gin.Context)
}

type provinsiController struct {
	jwtService service.JWTService
	provinsiService service.ProvinsiService
}

func NewProvinsiController(us service.ProvinsiService) ProvinsiController {
	return &provinsiController{
		provinsiService: us,
	}
}

func(uc *provinsiController) GetAllProvinsi(ctx *gin.Context) {
	result, err := uc.provinsiService.GetAllProvinsi(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Provinsi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Provinsi", result)
	ctx.JSON(http.StatusOK, res)
}