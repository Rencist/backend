package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
)

type SeriController interface {
	CreateSeri(ctx *gin.Context)
	GetAllSeri(ctx *gin.Context)
	DeleteSeri(ctx *gin.Context)
	UpdateSeri(ctx *gin.Context)
	FindSeriByID(ctx *gin.Context)
	UpsertRating(ctx *gin.Context)
	GetRating(ctx *gin.Context)
	GetRatingSeri(ctx *gin.Context)
}

type seriController struct {
	jwtService service.JWTService
	seriService service.SeriService
}

func NewSeriController(us service.SeriService, js service.JWTService) SeriController {
	return &seriController{
		seriService: us,
		jwtService:js,
	}
}

func(uc *seriController) CreateSeri(ctx *gin.Context) {
	var seri dto.SeriCreateDTO
	err := ctx.ShouldBind(&seri)
	result, err := uc.seriService.CreateSeri(ctx.Request.Context(), seri)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Seri", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *seriController) GetAllSeri(ctx *gin.Context) {
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

	if sort != "judul" && sort != "skor" && sort != "total_pembaca" {
		sort = "total_pembaca"
	}

	if sort == "judul" {
		sort = sort + " asc"
	} else {
		sort = sort + " desc"
	}

	fmt.Println("sort: ", sort)
		
	result, err := uc.seriService.GetAllSeri(ctx.Request.Context(), pagination, filter, search, sort)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Seri", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *seriController) DeleteSeri(ctx *gin.Context) {
	seriID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	err = uc.seriService.DeleteSeri(ctx.Request.Context(), seriID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Seri", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func(uc *seriController) UpdateSeri(ctx *gin.Context) {
	seriID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	var seri dto.SeriUpdateDTO
	seri.ID = seriID
	err = ctx.ShouldBind(&seri)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = uc.seriService.UpdateSeri(ctx.Request.Context(), seri)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Seri", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func(uc *seriController) FindSeriByID(ctx *gin.Context) {
	seriID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Detail Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := uc.seriService.FindSeriByID(ctx.Request.Context(), seriID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Detail Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mendapatkan Detail Seri", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *seriController) UpsertRating(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var rating dto.RatingCreateDTO
	err = ctx.ShouldBind(&rating)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Memberi Rating Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = uc.seriService.UpsertRating(ctx.Request.Context(), rating.SeriID, rating.Rating, userID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Memberi Rating Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Memberi Rating Seri", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func(uc *seriController) GetRating(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	result, err := uc.seriService.FindUserRating(ctx.Request.Context(), userID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Detail Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mendapatkan Detail Seri", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *seriController) GetRatingSeri(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	seriID, err := strconv.Atoi(ctx.Param("seri_id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := uc.seriService.FindUserSeriRating(ctx.Request.Context(), userID, seriID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Detail Seri", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mendapatkan Detail Seri", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *seriController) QueryArrayRequest(ctx *gin.Context, key string) ([]map[string]string){
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

