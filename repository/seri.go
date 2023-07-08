package repository

import (
	"context"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeriRepository interface {
	GetTotalData(ctx context.Context, search string) (int64, error)
	CreateSeri(ctx context.Context, seri entity.Seri) (entity.Seri, error)
	GetAllSeri(ctx context.Context, pagination entity.Pagination, filter []int, search string, sort string) (dto.PaginationResponse, error)
	FindSeriByIDDTOResponse(ctx context.Context, seriID int) (dto.SeriResponseDTO, error)
	DeleteSeri(ctx context.Context, seriID int) (error)
	UpdateSeri(ctx context.Context, seri entity.Seri) (error)
	AddRating(ctx context.Context, seriID int, rating float32, userID uuid.UUID) (error)
	CheckRatingUser(ctx context.Context, seriID int, userID uuid.UUID) (bool, error)
	UpdateRating(ctx context.Context, seriID int, rating float32, userID uuid.UUID) (error)
	FindSeryByID(ctx context.Context, seriID int) (entity.Seri, error)
	FindPenulisBySeriID(ctx context.Context, seriID int) ([]entity.Penulis, error)
	GetSeriIDGenreByGenreID(ctx context.Context, search string, GenreID []int) ([]int64, error)
	FindUserRating(ctx context.Context, userID uuid.UUID) ([]dto.RatingReponse, error)
	FindUserSeriRating(ctx context.Context, userID uuid.UUID, seriID int) (dto.RatingReponse, error)
	GetIDSeriTerakhir(ctx context.Context) (int, error)
}

type seriConnection struct {
	connection *gorm.DB
}

func NewSeriRepository(db *gorm.DB) SeriRepository {
	return &seriConnection{
		connection: db,
	}
}

func (db *seriConnection) GetTotalData(ctx context.Context, search string) (int64, error) {
	var totalData int64
	bc := db.connection.Model(&entity.Seri{})
	if search != "" {
		bc.Where("LOWER(judul) LIKE LOWER(?)", "%"+search+"%")
	}
	bc.Count(&totalData)
	if bc.Error != nil {
		return 0, bc.Error
	}
	return totalData, nil
}

func(db *seriConnection) CreateSeri(ctx context.Context, seri entity.Seri) (entity.Seri, error) {
	uc := db.connection.Create(&seri)
	if uc.Error != nil {
		return entity.Seri{}, uc.Error
	}
	return seri, nil
}

func(db *seriConnection) GetAllSeri(ctx context.Context, pagination entity.Pagination, filter []int, search string, sort string) (dto.PaginationResponse, error) {
	var listSeri []entity.Seri
	var totalData int64
	var tx *gorm.DB
	if len(filter) > 0 {
		seriGenre, err := db.GetSeriIDGenreByGenreID(ctx, search, filter)
		if err != nil {
			return dto.PaginationResponse{}, err
		}
		totalData = int64(len(seriGenre))
		tx = db.connection.Debug().Scopes(common.Pagination(&pagination, totalData)).Preload("Mangas").Preload("SeriGenre").Preload("PenulisSeri").Where("id IN ?", seriGenre)
		if tx.Error != nil {
			return dto.PaginationResponse{}, tx.Error
		}
	} else {
		totalData, _ = db.GetTotalData(ctx, search)
		tx = db.connection.Debug().Scopes(common.Pagination(&pagination, totalData)).Preload("Mangas").Preload("SeriGenre").Preload("PenulisSeri")
		if tx.Error != nil {
			return dto.PaginationResponse{}, tx.Error
		}
	}
	
	if sort != "" {
		tx = tx.Order(sort)
		if tx.Error != nil {
			return dto.PaginationResponse{}, tx.Error
		}
	}

	if search != "" {
		tx = tx.Where("LOWER(judul) LIKE LOWER(?)", "%"+search+"%")
		if tx.Error != nil {
			return dto.PaginationResponse{}, tx.Error
		}
	}
	
	tx.Find(&listSeri)
	var listSeriDTOArray []dto.SeriResponseDTO
	for _, res := range listSeri {
		var penerbit entity.Penerbit
		db.connection.Where("id = ?", res.PenerbitID).Take(&penerbit)
		var listSeriDTO dto.SeriResponseDTO
		listSeriDTO.ID = res.ID
		listSeriDTO.Judul = res.Judul
		listSeriDTO.Sinopsis = res.Sinopsis
		listSeriDTO.TahunTerbit = res.TahunTerbit
		listSeriDTO.Skor = res.Skor
		listSeriDTO.TotalPenilai = res.TotalPenilai
		listSeriDTO.TotalPembaca = res.TotalPembaca
		listSeriDTO.PenerbitID = res.PenerbitID
		listSeriDTO.NamaPenerbit = penerbit.Nama
		listSeriDTO.Foto = res.Foto
		for _, res := range res.Mangas {
			listSeriDTO.Manga = append(listSeriDTO.Manga, res)
		}
		for _, res := range res.PenulisSeri {
			var penulis entity.Penulis
			db.connection.Where("id = ?", res.PenulisID).Take(&penulis)
			listSeriDTO.Penulis = append(listSeriDTO.Penulis, penulis)
		}
		for _, res := range res.SeriGenre {
			var genre entity.Genre
			db.connection.Where("id = ?", res.GenreID).Take(&genre)
			listSeriDTO.Genre = append(listSeriDTO.Genre, genre)
		}
		
		listSeriDTOArray = append(listSeriDTOArray, listSeriDTO)
	}

	meta := dto.Meta{
		Page: pagination.Page,
		PerPage: pagination.PerPage,
		MaxPage: pagination.MaxPage,
		TotalData: totalData,
	}
	paginationResponse := dto.PaginationResponse{
		DataPerPage: listSeriDTOArray,
		Meta: meta,
	}
	return paginationResponse, nil
}

func(db *seriConnection) FindSeriByIDDTOResponse(ctx context.Context, seriID int) (dto.SeriResponseDTO, error) {
	var seri entity.Seri
	ux := db.connection.Where("id = ?", seriID).Preload("Mangas").Preload("SeriGenre").Preload("PenulisSeri").Take(&seri)
	var penerbit entity.Penerbit
	db.connection.Where("id = ?", seri.PenerbitID).Take(&penerbit)
	var seriDTO dto.SeriResponseDTO
	seriDTO.ID = seri.ID
	seriDTO.Judul = seri.Judul
	seriDTO.Sinopsis = seri.Sinopsis
	seriDTO.TahunTerbit = seri.TahunTerbit
	seriDTO.Skor = seri.Skor
	seriDTO.TotalPenilai = seri.TotalPenilai
	seriDTO.TotalPembaca = seri.TotalPembaca
	seriDTO.PenerbitID = seri.PenerbitID
	seriDTO.NamaPenerbit = penerbit.Nama
	seriDTO.Foto = seri.Foto
	for _, res := range seri.Mangas {
		seriDTO.Manga = append(seriDTO.Manga, res)
	}
	for _, res := range seri.PenulisSeri {
		var penulis entity.Penulis
		db.connection.Where("id = ?", res.PenulisID).Take(&penulis)
		seriDTO.Penulis = append(seriDTO.Penulis, penulis)
	}
	for _, res := range seri.SeriGenre {
		var genre entity.Genre
		db.connection.Where("id = ?", res.GenreID).Take(&genre)
		seriDTO.Genre = append(seriDTO.Genre, genre)
	}
	if ux.Error != nil {
		return seriDTO, ux.Error
	}
	return seriDTO, nil
}

func(db *seriConnection) DeleteSeri(ctx context.Context, seriID int) (error) {
	uc := db.connection.Delete(&entity.Seri{}, &seriID)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func(db *seriConnection) UpdateSeri(ctx context.Context, seri entity.Seri) (error) {
	uc := db.connection.Updates(&seri)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func(db *seriConnection) AddRating(ctx context.Context, seriID int, rating float32, userID uuid.UUID) (error) {
	uc := db.connection.Model(&entity.Seri{}).Where("id = ?", seriID).Update("total_penilai", gorm.Expr("total_penilai + 1"))
	if uc.Error != nil {
		return uc.Error
	}
	uc = db.connection.Model(&entity.Seri{}).Where("id = ?", seriID).Update("skor", gorm.Expr("round((((total_penilai - 1) * skor) + ?) / total_penilai, 2)", rating))
	if uc.Error != nil {
		return uc.Error
	}

	ratingID := uuid.New()
	entityRating := entity.Rating{
		ID: ratingID,
		Rating: rating,
		SeriID: seriID,
		UserID: userID,
	}
	uc = db.connection.Create(&entityRating)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func(db *seriConnection) CheckRatingUser(ctx context.Context, seriID int, userID uuid.UUID) (bool, error) {
	var rating entity.Rating
	uc := db.connection.Where("seri_id = ? AND user_id = ?", seriID, userID).Take(&rating)
	if uc.Error != nil {
		return false, uc.Error
	}
	if rating.ID != uuid.Nil {
		return true, nil
	}
	return false, nil
}

func(db *seriConnection) UpdateRating(ctx context.Context, seriID int, rating float32, userID uuid.UUID) (error) {
	var entityRating entity.Rating
	uc := db.connection.Where("seri_id = ? AND user_id = ?", seriID, userID).Take(&entityRating)
	if uc.Error != nil {
		return uc.Error
	}

	newRating := entityRating.Rating - rating
	uc = db.connection.Model(&entity.Seri{}).Where("id = ?", seriID).Update("skor", gorm.Expr("round(((skor * total_penilai) + ?) / total_penilai, 2)", newRating))
	if uc.Error != nil {
		return uc.Error
	}

	uc = db.connection.Model(&entity.Rating{}).Where("seri_id = ? AND user_id = ?", seriID, userID).Update("rating", rating)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func(db *seriConnection) FindSeryByID(ctx context.Context, seriID int) (entity.Seri, error) {
	var seri entity.Seri
	ux := db.connection.Where("id = ?", seriID).Find(&seri)
	if ux.Error != nil {
		return seri, ux.Error
	}
	return seri, nil
}

func(db *seriConnection) FindPenulisBySeriID(ctx context.Context, seriID int) ([]entity.Penulis, error) {
	var penulisSeri []entity.PenulisSeri
	var penulis []entity.Penulis
	ux := db.connection.Where("seri_id = ?", seriID).Find(&penulisSeri)
	if ux.Error != nil {
		return nil, ux.Error
	}
	for _, res := range penulisSeri {
		var penulisEntity entity.Penulis
		ux := db.connection.Where("id = ?", res.PenulisID).Find(&penulisEntity)
		if ux.Error != nil {
			return nil, ux.Error
		}
		penulis = append(penulis, penulisEntity)
	}
	return penulis, nil
}

func(db *seriConnection) GetSeriIDGenreByGenreID(ctx context.Context, search string, GenreID []int) ([]int64, error) {
	var genreIDFilter []int64
	search = "%" + search + "%"
	ux := db.connection.Raw(`
		select kuda.manga_id
		from
			(
				select
					count(1) as jumlah_manga_genre,
					sapi.manga_id
				from
					(
						select
							s.genre_id,
							m.judul as manga,
							m.id as manga_id,
							g.nama as genre
						from
							seri_genres s
							JOIN genres g ON s.genre_id = g.id
							JOIN seris m ON s.seri_id = m.id
						where
							s.genre_id IN (?)
							and 
							LOWER(judul) LIKE LOWER(?)
					) sapi
				group by(sapi.manga_id)
			) kuda
		where
		kuda.jumlah_manga_genre > ?`, GenreID, search, len(GenreID) - 1).Scan(&genreIDFilter)
	if ux.Error != nil {
		return nil, ux.Error
	}
	return genreIDFilter, nil
}

func(db *seriConnection) FindUserRating(ctx context.Context, userID uuid.UUID) ([]dto.RatingReponse, error) {
	var rating []entity.Rating
	ux := db.connection.Where("user_id = ?", userID).Find(&rating)
	if ux.Error != nil {
		return []dto.RatingReponse{}, ux.Error
	}
	var ratingDTOArray []dto.RatingReponse
	for _, res := range rating {
		ratingDTO := dto.RatingReponse{
			ID: res.ID,
			Rating: res.Rating,
			SeriID: res.SeriID,
			UserID: res.UserID,
		}
		ratingDTOArray = append(ratingDTOArray, ratingDTO)
	}
	
	return ratingDTOArray, nil
}

func(db *seriConnection) FindUserSeriRating(ctx context.Context, userID uuid.UUID, seriID int) (dto.RatingReponse, error) {
	var rating entity.Rating
	ux := db.connection.Where("user_id = ? AND seri_id = ?", userID, seriID).Take(&rating)
	if ux.Error != nil {
		return dto.RatingReponse{}, ux.Error
	}
	ratingDTO := dto.RatingReponse{
		ID: rating.ID,
		Rating: rating.Rating,
		SeriID: rating.SeriID,
		UserID: rating.UserID,
	}
	return ratingDTO, nil
}

func(db *seriConnection) GetIDSeriTerakhir(ctx context.Context) (int, error) {
	var seri entity.Seri
	ux := db.connection.Order("id desc").First(&seri)
	if ux.Error != nil {
		return 0, ux.Error
	}
	return seri.ID, nil
}