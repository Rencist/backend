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
	GetUserTransaction(ctx context.Context, userID uuid.UUID) ([]dto.TransactionResponse, error)
}

type movieService struct {
	movieRepository repository.MovieRepository
	userRepository repository.UserRepository
}

func NewMovieService(ur repository.MovieRepository, urr repository.UserRepository) MovieService {
	return &movieService{
		movieRepository: ur,
		userRepository: urr,
	}
}

func(us *movieService) GetAllMovie(ctx context.Context, pagination entity.Pagination, filter []int, search string, sort string) (dto.PaginationResponse, error) {
	return us.movieRepository.GetAllMovie(ctx, pagination, filter, search, sort)
}

func(us *movieService) GetMovieByID(ctx context.Context, movieID int) (entity.Movie, error) {
	return us.movieRepository.FindMovieByID(ctx, movieID)
}

func(us *movieService) CreateTransaction(ctx context.Context, transaction dto.TransactionCreateDTO) (entity.Transaction, error) {
	user, _ := us.userRepository.FindUserByID(ctx, transaction.UserID)
	movieID, _ := strconv.Atoi(transaction.MovieID)
	totalPrice, _ := strconv.Atoi(transaction.TotalPrice)
	movie, _ := us.movieRepository.FindMovieByID(ctx, movieID)
	if user.Balance < totalPrice {
		return entity.Transaction{}, errors.New("Saldo Anda Tidak Cukup")
	}
	movieAge, _ := strconv.Atoi(movie.AgeRating)
	if user.Age < movieAge {
		return entity.Transaction{}, errors.New("Umur Anda Tidak Cukup")
	}
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
	us.userRepository.WithdrawBalance(ctx, user.ID, totalPrice)
	return res, err
}

func(us *movieService) GetAvailableSeat(ctx context.Context, movieID int) ([]dto.AvalilableSeat, error) {
	return us.movieRepository.GetAvailableSeat(ctx, movieID)
}

func(us *movieService) GetUserTransaction(ctx context.Context, userID uuid.UUID) ([]dto.TransactionResponse, error) {
	listTransaction, err := us.movieRepository.GetUserTransaction(ctx, userID)
	if err != nil {
		return nil, err
	}
	responseDTO := []dto.TransactionResponse{}
	for _, v := range listTransaction {
		listSeat, _ := us.movieRepository.FindSeatByTransactionID(ctx, v.ID)
		seatDTO := []dto.SeatCreateDTO{}
		for _, w := range listSeat {
			seatDTO = append(seatDTO, dto.SeatCreateDTO{
				Seat: w.Seat,
			})
		}
		movieName, _ := us.movieRepository.FindMovieByID(ctx, v.MovieID)
		responseDTO = append(responseDTO, dto.TransactionResponse{
			ID: v.ID,
			MovieName: movieName.Title,
			TotalPrice: strconv.Itoa(v.TotalPrice),
			Seat: seatDTO,
		})
	}
	return responseDTO, nil
}
