package service

import (
	"context"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"
)

type MovieService interface {
	GetAllMovie(ctx context.Context) ([]entity.Movie, error)
	GetMovieByID(ctx context.Context, movieID int) (entity.Movie, error)
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(ur repository.MovieRepository) MovieService {
	return &movieService{
		movieRepository: ur,
	}
}

func(us *movieService) GetAllMovie(ctx context.Context) ([]entity.Movie, error) {
	return us.movieRepository.GetAllMovie(ctx)
}

func(us *movieService) GetMovieByID(ctx context.Context, movieID int) (entity.Movie, error) {
	return us.movieRepository.FindMovieByID(ctx, movieID)
}