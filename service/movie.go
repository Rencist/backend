package service

import (
	"context"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"
)

type MovieService interface {
	GetAllMovie(ctx context.Context, pagination entity.Pagination, filter []int, search string, sort string) (dto.PaginationResponse, error)
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

func(us *movieService) GetAllMovie(ctx context.Context, pagination entity.Pagination, filter []int, search string, sort string) (dto.PaginationResponse, error) {
	return us.movieRepository.GetAllMovie(ctx, pagination, filter, search, sort)
}

func(us *movieService) GetMovieByID(ctx context.Context, movieID int) (entity.Movie, error) {
	return us.movieRepository.FindMovieByID(ctx, movieID)
}