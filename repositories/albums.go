package repositories

import (
	"myproject/database"
	"myproject/models"

	"gorm.io/gorm"
)

type AlbumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository() *AlbumRepository {
	return &AlbumRepository{db: database.GetDb()}
}

func (a *AlbumRepository) GetAllAlbums() ([]models.Album, error) {
	var albums []models.Album
	err := a.db.Find(&albums).Error
	return albums, err
}

func (a *AlbumRepository) GetAlbumByID(id int) (*models.Album, error) {
	var album models.Album
	err := a.db.First(&album, id).Error
	return &album, err
}

func (a *AlbumRepository) CreateAlbum(album *models.Album) error {
	err := a.db.Create(&album).Error
	return err
}

func (a *AlbumRepository) UpdateAlbum(album *models.Album) error {
	err := a.db.Save(&album).Error
	return err
}

func (a *AlbumRepository) DeleteAlbumByID(album *models.Album) error {
	err := a.db.Delete(&album).Error
	return err
}
