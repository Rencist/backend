package controller

import (
	"net/http"
	"strconv"
	"strings"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MovieController interface {
	GetAllMovie(ctx *gin.Context)
	GetMovieByID(ctx *gin.Context)
	CreateTransaction(ctx *gin.Context)
	GetAvailableSeat(ctx *gin.Context)
	GetUserTransaction(ctx *gin.Context)
	DeleteTransaction(ctx *gin.Context)
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

func(uc *movieController) CreateTransaction(ctx *gin.Context) {
	ctx.PostFormArray("seat")
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	
	var transaction dto.TransactionCreateDTO
	transaction.TotalPrice = ctx.Request.PostForm.Get("total_price")
	transaction.MovieID = ctx.Request.PostForm.Get("movie_id")
	total_seat, _ := strconv.Atoi(ctx.Request.PostForm.Get("total_seat"))
	if total_seat <= 0 && total_seat > 6 {
		res := common.BuildErrorResponse("Gagal Menambahkan Transaksi", "Total Seat Tidak Memenuhi Syarat", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	for i := 0; i < total_seat; i++ {
		var seat dto.SeatCreateDTO
		seat_number, _ := strconv.Atoi(ctx.Request.PostForm.Get("seat["+strconv.Itoa(i)+"][seat]"))
		seat.Seat = seat_number
		transaction.Seat = append(transaction.Seat, seat)
	}
	transaction.UserID = userID
	_, err = uc.movieService.CreateTransaction(ctx.Request.Context(), transaction)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Transaksi", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func(uc *movieController) GetAvailableSeat(ctx *gin.Context) {
	movieID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Detail Movie", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := uc.movieService.GetAvailableSeat(ctx.Request.Context(), movieID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Movie", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan Movie", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *movieController) GetUserTransaction(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}


	result, err := uc.movieService.GetUserTransaction(ctx.Request.Context(), userID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Transaksi", result)
	ctx.JSON(http.StatusOK, res)
}

func(uc *movieController) DeleteTransaction(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var transactionDTO dto.DeleteTransactionDTO
	err = ctx.ShouldBind(&transactionDTO)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	transactionID, err := uuid.Parse(transactionDTO.TransactionID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = uc.movieService.DeleteTransaction(ctx.Request.Context(), transactionID, userID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menghapus Transaksi", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}