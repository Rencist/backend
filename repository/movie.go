package repository

import (
	"context"
	"tamiyochi-backend/entity"

	"gorm.io/gorm"
)

type MovieRepository interface {
	GetAllMovie(ctx context.Context) ([]entity.Movie, error)
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


func(db *movieConnection) GetAllMovie(ctx context.Context) ([]entity.Movie, error) {
	var listMovie []entity.Movie
	tx := db.connection.Find(&listMovie)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listMovie, nil
}

func(db *movieConnection) FindMovieByID(ctx context.Context, movieID int) (entity.Movie, error) {
	var movie entity.Movie
	ux := db.connection.Where("id = ?", movieID).Take(&movie)
	if ux.Error != nil {
		return movie, ux.Error
	}
	return movie, nil
}