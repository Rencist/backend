package controller

import (
	"net/http"
	"strconv"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

type CartController interface {
	CreateCart(ctx *gin.Context)
	FindCartByUserID(ctx *gin.Context)
	DeleteCart(ctx *gin.Context)
	DeleteAllByMangaIDCart(ctx *gin.Context)
}

type cartController struct {
	jwtService service.JWTService
	cartService service.CartService
}

func NewCartController(us service.CartService, jwts service.JWTService) CartController {
	return &cartController{
		cartService: us,
		jwtService: jwts,
	}
}

func(uc *cartController) CreateCart(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	
	var cart dto.CartCreateDTO
	err = ctx.ShouldBind(&cart)
	cart.UserID = userID
	result, err := uc.cartService.CreateCart(ctx.Request.Context(), cart)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Cart", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Cart", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *cartController) FindCartByUserID(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	result, err := uc.cartService.FindCartByUserID(ctx.Request.Context(), userID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Cart", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mendapatkan List Cart", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *cartController) DeleteCart(ctx *gin.Context) {
	mangaID, err := strconv.Atoi(ctx.Param("manga_id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Cart", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	err = uc.cartService.DeleteCart(ctx.Request.Context(), mangaID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Cart", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Cart", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func(uc *cartController) DeleteAllByMangaIDCart(ctx *gin.Context) {
	mangaID, err := strconv.Atoi(ctx.Param("manga_id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Cart", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	err = uc.cartService.DeleteAllByMangaIDCart(ctx.Request.Context(), mangaID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Cart", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Cart", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
