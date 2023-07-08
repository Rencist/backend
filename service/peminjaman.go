package service

import (
	"context"
	"errors"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"
	"time"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type PeminjamanService interface {
	CreatePeminjaman(ctx context.Context, peminjamanDTO dto.PeminjamanCreateDTO) (entity.Peminjaman, error)
	GetAllPeminjamanUser(ctx context.Context, userID uuid.UUID) ([]dto.PeminjamanReponse, error)
	PaidDenda(ctx context.Context, UserID uuid.UUID, denda dto.DendaCreateDTO) (error)
}

type peminjamanService struct {
	peminjamanRepository repository.PeminjamanRepository
	cartRepository repository.CartRepository
	mangaRepository repository.MangaRepository
	seriRepository repository.SeriRepository
}

func NewPeminjamanService(ur repository.PeminjamanRepository, cr repository.CartRepository, mr repository.MangaRepository, sr repository.SeriRepository) PeminjamanService {
	return &peminjamanService{
		peminjamanRepository: ur,
		cartRepository: cr,
		mangaRepository: mr,
		seriRepository: sr,
	}
}

// Menunggu Konfirmasi, Belum Diambil, Sedang Dipinjam, Sudah Dikembalikan
func(us *peminjamanService) CreatePeminjaman(ctx context.Context, peminjamanDTO dto.PeminjamanCreateDTO) (entity.Peminjaman, error) {
	cekCart, err := us.cartRepository.CekTotalCart(ctx, peminjamanDTO.UserID)
	if err != nil {
		return entity.Peminjaman{}, err
	}
	if cekCart == 0 {
		return entity.Peminjaman{}, errors.New("Cart Kosong")
	}
	peminjaman := entity.Peminjaman{}
	err = smapping.FillStruct(&peminjaman, smapping.MapFields(peminjamanDTO))
	if err != nil {
		return peminjaman, err
	}
	peminjaman.ID = uuid.New()
	peminjaman.StatusPeminjaman = "Menunggu Konfirmasi"
	peminjaman.BatasPengembalian = time.Now().AddDate(0, 0, 7)
	peminjaman.TanggalPengembalian = time.Time{}
	createPeminjaman, err := us.peminjamanRepository.CreatePeminjaman(ctx, peminjaman)
	if err != nil {
		return peminjaman, err
	}

	cart, err := us.cartRepository.FindCartByUserID(ctx, peminjamanDTO.UserID)
	if err != nil {
		return peminjaman, err
	}

	for _, res := range cart {
		peminjamanManga := entity.PeminjamanManga{}
		peminjamanManga.ID = uuid.New()
		peminjamanManga.MangaID = res.MangaID
		peminjamanManga.PeminjamanID = createPeminjaman.ID
		_, err = us.peminjamanRepository.CreatePeminjamanManga(ctx, peminjamanManga)
		if err != nil {
			return peminjaman, err
		}
	}

	err = us.cartRepository.DeleteAllUserCart(ctx, peminjamanDTO.UserID)
	if err != nil {
		return peminjaman, err
	}
	
	return createPeminjaman, err
}

func(us *peminjamanService) GetAllPeminjamanUser(ctx context.Context, userID uuid.UUID) ([]dto.PeminjamanReponse, error) {
	var peminjaman []entity.Peminjaman
	var listPeminjamanDTOArray []dto.PeminjamanReponse
	var listPeminjamanDTO dto.PeminjamanReponse
	var listPeminjamanManga []entity.PeminjamanManga

	peminjaman, err := us.peminjamanRepository.GetAllPeminjamanUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, res := range peminjaman {
		listPeminjamanManga, err = us.peminjamanRepository.FindPeminjamanMangaByPeminjamanID(ctx, res.ID)
		for _, res2 := range listPeminjamanManga {
			var manga entity.Manga
			manga, err = us.mangaRepository.FindMangaByID(ctx, res2.MangaID)
			if err != nil {
				return nil, err
			}

			var seri entity.Seri
			seri, err = us.seriRepository.FindSeryByID(ctx, manga.SeriID)
			if err != nil {
				return nil, err
			}

			var penulis []entity.Penulis
			penulis, err = us.seriRepository.FindPenulisBySeriID(ctx, seri.ID)
			if err != nil {
				return nil, err
			}

			var denda entity.Denda
			denda, err = us.peminjamanRepository.FindDendaByPeminjamanID(ctx, res.ID)
			if err != nil {
				return nil, err
			}

			listPeminjamanDTO.IDPeminjaman = res.ID
			listPeminjamanDTO.IDPeminjamanManga = res2.ID
			listPeminjamanDTO.IDSeri = seri.ID
			listPeminjamanDTO.IDManga = manga.ID
			listPeminjamanDTO.IDSeri = manga.SeriID
			listPeminjamanDTO.IDDenda = denda.ID

			listPeminjamanDTO.Volume = manga.Volume
			listPeminjamanDTO.Judul = seri.Judul
			listPeminjamanDTO.Foto = seri.Foto
			listPeminjamanDTO.TanggalPeminjaman = res.CreatedAt
			listPeminjamanDTO.BatasPengembalian = res.BatasPengembalian
			listPeminjamanDTO.StatusPeminjaman = res.StatusPeminjaman
			listPeminjamanDTO.Penulis = penulis
			listPeminjamanDTO.Denda = denda.TotalDenda
			listPeminjamanDTO.IsDendaLunas = denda.IsLunas

			listPeminjamanDTOArray = append(listPeminjamanDTOArray, listPeminjamanDTO)
		}
	}
	
	return listPeminjamanDTOArray, err
}

func(us *peminjamanService) PaidDenda(ctx context.Context, UserID uuid.UUID, denda dto.DendaCreateDTO) (error) {
	err := us.peminjamanRepository.PaidDenda(ctx, UserID, denda.PeminjamanID)
	if err != nil {
		return err
	}
	return nil
}