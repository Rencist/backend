package service

import (
	"context"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"

	"github.com/google/uuid"
)

type CartService interface {
	CreateCart(ctx context.Context, cartDTO dto.CartCreateDTO) (entity.Cart, error)
	FindCartByUserID(ctx context.Context, userID uuid.UUID) (dto.CartResponse, error)
	DeleteCart(ctx context.Context, mangaID int) (error)
	DeleteAllByMangaIDCart(ctx context.Context, mangaID int) (error)
}

type cartService struct {
	cartRepository repository.CartRepository
}

func NewCartService(ur repository.CartRepository) CartService {
	return &cartService{
		cartRepository: ur,
	}
}

func(us *cartService) CreateCart(ctx context.Context, cartDTO dto.CartCreateDTO) (entity.Cart, error) {
	// cart := entity.Cart{}
	// err := smapping.FillStruct(&cart, smapping.MapFields(cartDTO))
	// if err != nil {
	// 	return cart, err
	// }
	return us.cartRepository.CreateCart(ctx, cartDTO)
}

func(us *cartService) FindCartByUserID(ctx context.Context, userID uuid.UUID) (dto.CartResponse, error) {
	// res, err := us.cartRepository.FindCartByUserID(ctx, userID)
	// if err != nil {
	// 	return nil, err
	// }
	// for i, _ := range res {
	// }
	return us.cartRepository.FindCartByUserIDResponse(ctx, userID)
}

func(us *cartService) DeleteCart(ctx context.Context, mangaID int) (error) {
	return us.cartRepository.DeleteCart(ctx, mangaID)
}

func(us *cartService) DeleteAllByMangaIDCart(ctx context.Context, mangaID int) (error) {
	return us.cartRepository.DeleteAllByMangaIDCart(ctx, mangaID)
}