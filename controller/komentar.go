package controller

import (
	"net/http"
	"strconv"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

type KomentarController interface {
	CreateKomentar(ctx *gin.Context)
	FindKomentarBySeriID(ctx *gin.Context)
}

type komentarController struct {
	jwtService service.JWTService
	komentarService service.KomentarService
}

func NewKomentarController(us service.KomentarService, jwts service.JWTService) KomentarController {
	return &komentarController{
		komentarService: us,
		jwtService: jwts,
	}
}

func(uc *komentarController) CreateKomentar(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var komentar dto.KomentarCreateDTO
	err = ctx.ShouldBind(&komentar)
	komentar.UserID = userID
	result, err := uc.komentarService.CreateKomentar(ctx.Request.Context(), komentar)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Komentar", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Komentar", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *komentarController) FindKomentarBySeriID(ctx *gin.Context) {
	var pagination entity.Pagination
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page <= 0 {
		page = 1
	}
	pagination.Page = page

	perPage, _ := strconv.Atoi(ctx.Query("per_page"))
	if perPage <= 0 {
		perPage = 5
	}
	pagination.PerPage = perPage
	
	komentarID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Detail Komentar", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := uc.komentarService.FindKomentarBySeriID(ctx.Request.Context(), komentarID, pagination)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Detail Komentar", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mendapatkan Detail Komentar", result)
	ctx.JSON(http.StatusOK, res)
}