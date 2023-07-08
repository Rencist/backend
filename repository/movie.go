package repository

import (
	"context"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"

	"gorm.io/gorm"
)

type MovieRepository interface {
	GetTotalData(ctx context.Context, search string) (int64, error)
	GetAllMovie(ctx context.Context, pagination entity.Pagination, filter []int, search string, sort string) (dto.PaginationResponse, error)
	FindMovieByID(ctx context.Context, movieID int) (entity.Movie, error)
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