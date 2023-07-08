package controller

import (
	"net/http"
	"strconv"
	"tamiyochi-backend/common"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

type MovieController interface {
	GetAllMovie(ctx *gin.Context)
	GetMovieByID(ctx *gin.Context)
}

type movieController struct {
	jwtService service.JWTService
	movieService service.MovieService
}

func NewMovieController(us service.MovieService, jwts service.JWTService) MovieController {
	return &movieController{
		movieService: us,
		jwtService: jwts,
	}
}

func(uc *movieController) GetAllMovie(ctx *gin.Context) {
	result, err := uc.movieService.GetAllMovie(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Movie", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Movie", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *movieController) GetMovieByID(ctx *gin.Context) {
	movieID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Detail Movie", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := uc.movieService.GetMovieByID(ctx.Request.Context(), movieID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Movie", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan Movie", result)
	ctx.JSON(http.StatusOK, res)
}
