package controller

import (
	"net/http"
	"strconv"
	"tamiyochi-backend/common"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

type KabupatenController interface {
	GetAllKabupaten(ctx *gin.Context)
	FindKabupatenByProvinsiID(ctx *gin.Context)
}

type kabupatenController struct {
	jwtService service.JWTService
	kabupatenService service.KabupatenService
}

func NewKabupatenController(us service.KabupatenService) KabupatenController {
	return &kabupatenController{
		kabupatenService: us,
	}
}

func(uc *kabupatenController) GetAllKabupaten(ctx *gin.Context) {
	result, err := uc.kabupatenService.GetAllKabupaten(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Kabupaten", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Kabupaten", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *kabupatenController) FindKabupatenByProvinsiID(ctx *gin.Context) {
	provinsiID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Kabupaten", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := uc.kabupatenService.FindKabupatenByProvinsiID(ctx.Request.Context(), provinsiID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Kabupaten", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Kabupaten", result)
	ctx.JSON(http.StatusOK, res)
}