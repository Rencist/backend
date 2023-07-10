package controller

import (
	"net/http"
	"strconv"
	"strings"
	"tamiyochi-backend/common"
	"tamiyochi-backend/entity"
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

	filterReq := uc.QueryArrayRequest(ctx, "filter")
	var filter[]int
	if len(filterReq) > 0 {
		for i := 0; i < len(filterReq[0]); i++ {
			filterToInt, _ := strconv.Atoi(filterReq[0][strconv.Itoa(i)])
			filter = append(filter, filterToInt)
		}
	}
	search := ctx.Query("search")
	sort := ctx.Query("sort")

	if sort != "title" && sort != "ticket_price" && sort != "age_rating" {
		sort = "title asc"
	}
	
	result, err := uc.movieService.GetAllMovie(ctx.Request.Context(), pagination, filter, search, sort)
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

func(uc *movieController) QueryArrayRequest(ctx *gin.Context, key string) ([]map[string]string){
	var dicts []map[string]string
	queryMap := ctx.Request.URL.Query()
	for k, v := range queryMap {
		if i:= strings.IndexByte(k, '['); i >= 1 && k[0:i] == key{
			if j := strings.IndexByte(k[i+1:], ']'); j >= 1{
				index, _ := strconv.Atoi(k[i+1: i+j+1])
				if index > len(dicts){
					ctx.JSON(200, gin.H{
						"403": "Check your data",
					})
					return nil
				}
				if index == len(dicts){
					tmp := make(map[string]string)
					dicts = append(dicts, tmp)
				}
				pre :=strings.IndexByte(k[i+j+2:], '[')
				last:=strings.IndexByte(k[i+j+2:], ']')
				dicts[index][k[i+j+3+pre: i+j+2+last]] = v[0]
			}
		}
	}
	return dicts
}