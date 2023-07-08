package controller

import (
	"net/http"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MangaController interface {
	CreateManga(ctx *gin.Context)
	GetAllManga(ctx *gin.Context)
	DeleteManga(ctx *gin.Context)
	UpdateManga(ctx *gin.Context)
}

type mangaController struct {
	jwtService service.JWTService
	mangaService service.MangaService
}

func NewMangaController(us service.MangaService) MangaController {
	return &mangaController{
		mangaService: us,
	}
}

func(uc *mangaController) CreateManga(ctx *gin.Context) {
	var manga dto.MangaCreateDTO
	err := ctx.ShouldBind(&manga)
	result, err := uc.mangaService.CreateManga(ctx.Request.Context(), manga)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Manga", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Manga", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *mangaController) GetAllManga(ctx *gin.Context) {
	result, err := uc.mangaService.GetAllManga(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Manga", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Manga", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *mangaController) DeleteManga(ctx *gin.Context) {
	mangaID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Manga", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	err = uc.mangaService.DeleteManga(ctx.Request.Context(), mangaID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Manga", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Manga", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func(uc *mangaController) UpdateManga(ctx *gin.Context) {
	mangaID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Manga", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	var manga dto.MangaUpdateDTO
	manga.ID = mangaID
	err = ctx.ShouldBind(&manga)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Manga", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = uc.mangaService.UpdateManga(ctx.Request.Context(), manga)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Manga", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Manga", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}