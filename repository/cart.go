package repository

import (
	"context"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(ctx context.Context, cartDTO dto.CartCreateDTO) (entity.Cart, error)
	FindCartByUserIDResponse(ctx context.Context, userID uuid.UUID) (dto.CartResponse, error)
	FindCartByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Cart, error)
	DeleteCart(ctx context.Context, mangaID int) error
	DeleteAllByMangaIDCart(ctx context.Context, mangaID int) error
	DeleteAllUserCart(ctx context.Context, userID uuid.UUID) error
	CekTotalCart(ctx context.Context, userID uuid.UUID) (int64, error)
}

type cartConnection struct {
	connection *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartConnection{
		connection: db,
	}
}

func (db *cartConnection) CreateCart(ctx context.Context, cartDTO dto.CartCreateDTO) (entity.Cart, error) {
	// var cartEntity entity.Cart
	cartEntity := entity.Cart{}
	err := smapping.FillStruct(&cartEntity, smapping.MapFields(cartDTO))
	if err != nil {
		return cartEntity, err
	}
	for i := 0; i < cartDTO.Jumlah; i++ {
		cartID := uuid.New()
		cartEntity.ID = cartID
		uc := db.connection.Create(&cartEntity)
		if uc.Error != nil {
			return entity.Cart{}, uc.Error
		}
	}
	return cartEntity, nil
}

func (db *cartConnection) FindCartByUserIDResponse(ctx context.Context, userID uuid.UUID) (dto.CartResponse, error) {
	var jumlahMangaCart dto.JumlahMangaCart
	var cart []entity.Cart

	var cartResponseDTO dto.CartResponse
	var cartDTO dto.Cart
	totalHarga := 0
	totalManga := 0
	ux := db.connection.Select("count(1) as jumlah_penyewa, manga_id").Group("manga_id").Find(&cart)
	for _, res := range cart {
		jumlahMangaCart.JumlahPenyewa = res.JumlahPenyewa
		jumlahMangaCart.MangaID = res.MangaID

		manga := entity.Manga{}
		ux := db.connection.Where("id = ?", res.MangaID).Find(&manga)
		if ux.Error != nil {
			return dto.CartResponse{}, ux.Error
		}
		seri := entity.Seri{}
		ux = db.connection.Where("id = ?", manga.SeriID).Find(&seri)
		if ux.Error != nil {
			return dto.CartResponse{}, ux.Error
		}
		cartDTO.ID = res.ID
		cartDTO.MangaID = manga.ID
		cartDTO.JumlahTersedia = manga.JumlahTersedia
		cartDTO.HargaSewa = manga.HargaSewa
		cartDTO.Volume = manga.Volume
		cartDTO.Foto = seri.Foto
		cartDTO.HargaSubTotal = manga.HargaSewa * res.JumlahPenyewa
		cartDTO.JumlahSewa = res.JumlahPenyewa
		cartDTO.JudulSeri = seri.Judul

		totalHarga += cartDTO.HargaSubTotal
		totalManga += res.JumlahPenyewa
		cartResponseDTO.Cart = append(cartResponseDTO.Cart, cartDTO)
	}
	cartResponseDTO.TotalHargaSewa = totalHarga
	cartResponseDTO.TotalPinjaman = totalManga

	if ux.Error != nil {
		return dto.CartResponse{}, ux.Error
	}
	return cartResponseDTO, nil
}

func (db *cartConnection) DeleteCart(ctx context.Context, mangaID int) error {
	var cart entity.Cart
	uc := db.connection.Where("manga_id = ?", mangaID).Limit(1).Take(&cart).Delete(&cart)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *cartConnection) DeleteAllByMangaIDCart(ctx context.Context, mangaID int) error {
	uc := db.connection.Where("manga_id = ?", mangaID).Delete(&entity.Cart{})
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *cartConnection) FindCartByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Cart, error) {
	var cart []entity.Cart
	ux := db.connection.Where("user_id = ?", userID).Find(&cart)
	if ux.Error != nil {
		return nil, ux.Error
	}
	return cart, nil
}

func (db *cartConnection) DeleteAllUserCart(ctx context.Context, userID uuid.UUID) error {
	uc := db.connection.Where("user_id = ?", userID).Delete(&entity.Cart{})
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *cartConnection) CekTotalCart(ctx context.Context, userID uuid.UUID) (int64, error) {
	var cartCount int64
	ux := db.connection.Model(&entity.Cart{}).Where("user_id = ?", userID).Count(&cartCount)
	if ux.Error != nil {
		return 0, ux.Error
	}
	return cartCount, nil
}
