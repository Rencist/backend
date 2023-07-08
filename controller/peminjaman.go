package controller

import (
	"net/http"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

type PeminjamanController interface {
	CreatePeminjaman(ctx *gin.Context)
	GetAllPeminjamanUser(ctx *gin.Context)
	PaidDenda(ctx *gin.Context)
}

type peminjamanController struct {
	jwtService service.JWTService
	peminjamanService service.PeminjamanService
}

func NewPeminjamanController(us service.PeminjamanService, jt service.JWTService) PeminjamanController {
	return &peminjamanController{
		peminjamanService: us,
		jwtService: jt,
	}
}

func(uc *peminjamanController) CreatePeminjaman(ctx *gin.Context) { 
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var peminjaman dto.PeminjamanCreateDTO
	peminjaman.UserID = userID
	err = ctx.ShouldBind(&peminjaman)
	result, err := uc.peminjamanService.CreatePeminjaman(ctx.Request.Context(), peminjaman)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Peminjaman", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Peminjaman", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *peminjamanController) GetAllPeminjamanUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	result, err := uc.peminjamanService.GetAllPeminjamanUser(ctx.Request.Context(), userID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Peminjaman", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Peminjaman", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *peminjamanController) PaidDenda(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	
	var denda dto.DendaCreateDTO
	err = ctx.ShouldBind(&denda)

	err = uc.peminjamanService.PaidDenda(ctx.Request.Context(), userID, denda)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Membayar Denda", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Membayar Denda", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}