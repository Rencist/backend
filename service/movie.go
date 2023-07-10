package service

import (
	"context"
	"errors"
	"strconv"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"

	"github.com/google/uuid"
)

type MovieService interface {
	GetAllMovie(ctx context.Context, pagination entity.Pagination, filter []int, search string, sort string) (dto.PaginationResponse, error)
	GetMovieByID(ctx context.Context, movieID int) (entity.Movie, error)
	CreateTransaction(ctx context.Context, transaction dto.TransactionCreateDTO) (entity.Transaction, error)
	GetAvailableSeat(ctx context.Context, movieID int) ([]dto.AvalilableSeat, error)
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(ur repository.MovieRepository) MovieService {
	return &movieService{
		movieRepository: ur,
	}
}

func(us *movieService) GetAllMovie(ctx context.Context, pagination entity.Pagination, filter []int, search string, sort string) (dto.PaginationResponse, error) {
	return us.movieRepository.GetAllMovie(ctx, pagination, filter, search, sort)
}

func(us *movieService) GetMovieByID(ctx context.Context, movieID int) (entity.Movie, error) {
	return us.movieRepository.FindMovieByID(ctx, movieID)
}

func(us *movieService) CreateTransaction(ctx context.Context, transaction dto.TransactionCreateDTO) (entity.Transaction, error) {
	movieID, _ := strconv.Atoi(transaction.MovieID)
	totalPrice, _ := strconv.Atoi(transaction.TotalPrice)
	transactionID := uuid.New()
	transactionEntity := entity.Transaction{
		ID: transactionID,
		UserID: transaction.UserID,
		MovieID: movieID,
		TotalPrice: totalPrice,
	}
	checkTransaction, _ := us.movieRepository.FindTransactionByMovieID(ctx, movieID)
	listSeat := []entity.Seat{}
	for _, v := range checkTransaction {
		checkSeat, _ := us.movieRepository.FindSeatByTransactionID(ctx, v.ID)
		for _, w := range checkSeat {
			listSeat = append(listSeat, w)
		}
	}
	for _, v := range transaction.Seat {
		for _, w := range listSeat {
			if v.Seat == w.Seat {
				return entity.Transaction{}, errors.New("Seat sudah terisi")
			}
		}
	}
	res, err := us.movieRepository.CreateTransaction(ctx, transactionEntity)
	for _, v := range transaction.Seat {
		seatEntity := entity.Seat{
			ID: uuid.New(),
			TransactionID: transactionID,
			Seat: v.Seat,
		}
		us.movieRepository.CreateSeat(ctx, seatEntity)
	}
	return res, err
}

func(us *movieService) GetAvailableSeat(ctx context.Context, movieID int) ([]dto.AvalilableSeat, error) {
	return us.movieRepository.GetAvailableSeat(ctx, movieID)
}
