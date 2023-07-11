package repository

import (
	"context"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieRepository interface {
	GetTotalData(ctx context.Context, search string) (int64, error)
	GetAllMovie(ctx context.Context, pagination entity.Pagination, filter []int, search string, sort string) (dto.PaginationResponse, error)
	FindMovieByID(ctx context.Context, movieID int) (entity.Movie, error)
	CreateTransaction(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error)
	CreateSeat(ctx context.Context, seat entity.Seat) (entity.Seat, error)
	FindTransactionByMovieID(ctx context.Context, movieID int) ([]entity.Transaction, error)
	GetAvailableSeat(ctx context.Context, movieID int) ([]dto.AvalilableSeat, error)
	FindSeatByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]entity.Seat, error)
	GetUserTransaction(ctx context.Context, userID uuid.UUID) ([]entity.Transaction, error)
	DeleteTransaction(ctx context.Context, transactionID uuid.UUID) error
	FindTransactionByID(ctx context.Context, transactionID uuid.UUID) (entity.Transaction, error)
}

type movieConnection struct {
	connection *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieConnection{
		connection: db,
	}
}

func (db *movieConnection) GetTotalData(ctx context.Context, search string) (int64, error) {
	var totalData int64
	bc := db.connection.Model(&entity.Movie{})
	if search != "" {
		bc.Where("LOWER(title) LIKE LOWER(?)", "%"+search+"%")
	}
	bc.Count(&totalData)
	if bc.Error != nil {
		return 0, bc.Error
	}
	return totalData, nil
}

func(db *movieConnection) GetAllMovie(ctx context.Context, pagination entity.Pagination, filter []int, search string, sort string) (dto.PaginationResponse, error) {
	var listMovie []entity.Movie
	totalData, err := db.GetTotalData(ctx, search)
	if err != nil {
		return dto.PaginationResponse{}, err
	}
	tx := db.connection.Debug().Scopes(common.Pagination(&pagination, totalData))
	if tx.Error != nil {
		return dto.PaginationResponse{}, tx.Error
	}

	if sort != "" {
		tx = tx.Order(sort)
		if tx.Error != nil {
			return dto.PaginationResponse{}, tx.Error
		}
	}

	if search != "" {
		tx = tx.Where("LOWER(title) LIKE LOWER(?)", "%"+search+"%")
		if tx.Error != nil {
			return dto.PaginationResponse{}, tx.Error
		}
	}

	tx.Find(&listMovie)
	var listMovieDTOArray []dto.MovieResponse
	for _, v := range listMovie {
		var listMovieDTO dto.MovieResponse
		listMovieDTO.ID = v.ID
		listMovieDTO.Title = v.Title
		listMovieDTO.Description = v.Description
		listMovieDTO.ReleaseDate = v.ReleaseDate
		listMovieDTO.PosterUrl = v.PosterUrl
		listMovieDTO.AgeRating = v.AgeRating
		listMovieDTO.TicketPrice = v.TicketPrice
		listMovieDTOArray = append(listMovieDTOArray, listMovieDTO)
	}

	meta := dto.Meta{
		Page: pagination.Page,
		PerPage: pagination.PerPage,
		MaxPage: pagination.MaxPage,
		TotalData: totalData,
	}
	paginationResponse := dto.PaginationResponse{
		DataPerPage: listMovieDTOArray,
		Meta: meta,
	}

	return paginationResponse, nil
}

func(db *movieConnection) FindMovieByID(ctx context.Context, movieID int) (entity.Movie, error) {
	var movie entity.Movie
	ux := db.connection.Where("id = ?", movieID).Take(&movie)
	if ux.Error != nil {
		return movie, ux.Error
	}
	return movie, nil
}

func(db *movieConnection) CreateTransaction(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	uc := db.connection.Create(&transaction)
	if uc.Error != nil {
		return entity.Transaction{}, uc.Error
	}
	return transaction, nil
}

func(db *movieConnection) CreateSeat(ctx context.Context, seat entity.Seat) (entity.Seat, error) {
	uc := db.connection.Create(&seat)
	if uc.Error != nil {
		return entity.Seat{}, uc.Error
	}
	return seat, nil
}

func(db *movieConnection) FindTransactionByMovieID(ctx context.Context, movieID int) ([]entity.Transaction, error) {
	var transaction []entity.Transaction
	ux := db.connection.Where("movie_id = ?", movieID, ).Find(&transaction)
	if ux.Error != nil {
		return transaction, ux.Error
	}
	return transaction, nil
}

func(db *movieConnection) GetAvailableSeat(ctx context.Context, movieID int) ([]dto.AvalilableSeat, error) {
	transaction, err := db.FindTransactionByMovieID(ctx, movieID)
	if err != nil {
		return nil, err
	}
	var seat []entity.Seat
	for _, v := range transaction {
		ux := db.connection.Where("transaction_id = ?", v.ID ).Find(&seat)
		if ux.Error != nil {
			return nil, ux.Error
		}
	}
	seatRes := []dto.AvalilableSeat{}
	for i := 1; i <= 64; i++ {
		lmao := false
		for _, v := range seat {
			if i == v.Seat {
				lmao = true
				break
			}
		}

		if lmao {
			continue
		}

		seatRes = append(seatRes, dto.AvalilableSeat{
			Seat: i,
		})
	}

	return seatRes, nil
}

func(db *movieConnection) FindSeatByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]entity.Seat, error) {
	var seat []entity.Seat
	ux := db.connection.Where("transaction_id = ?", transactionID, ).Find(&seat)
	if ux.Error != nil {
		return seat, ux.Error
	}
	return seat, nil
}

func(db *movieConnection) GetUserTransaction(ctx context.Context, userID uuid.UUID) ([]entity.Transaction, error) {
	var transaction []entity.Transaction
	ux := db.connection.Where("user_id = ?", userID, ).Find(&transaction)
	if ux.Error != nil {
		return transaction, ux.Error
	}
	return transaction, nil
}

func(db *movieConnection) DeleteTransaction(ctx context.Context, transactionID uuid.UUID) error {
	ux := db.connection.Where("id = ?", transactionID).Delete(&entity.Transaction{})
	if ux.Error != nil {
		return ux.Error
	}
	return nil
}

func(db *movieConnection) FindTransactionByID(ctx context.Context, transactionID uuid.UUID) (entity.Transaction, error) {
	var transaction entity.Transaction
	ux := db.connection.Where("id = ?", transactionID).Take(&transaction)
	if ux.Error != nil {
		return transaction, ux.Error
	}
	return transaction, nil
}
