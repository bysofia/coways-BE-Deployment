package repositories

import (
	"BackEnd/models"

	"gorm.io/gorm"
)

type SingerRepository interface {
	FindSingers() ([]models.Singer, error)
	GetSinger(ID int) (models.Singer, error)
	CreateSinger(singer models.Singer) (models.Singer, error)
	UpdateSinger(singer models.Singer) (models.Singer, error)
	DeleteSinger(singer models.Singer) (models.Singer, error)
}

func RepositorySinger(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindSingers() ([]models.Singer, error) {
	var singer []models.Singer
	err := r.db.Find(&singer).Error

	return singer, err
}

func (r *repository) GetSinger(ID int) (models.Singer, error) {
	var singer models.Singer
	err := r.db.First(&singer, ID).Error

	return singer, err
}

func (r *repository) CreateSinger(singer models.Singer) (models.Singer, error) {
	err := r.db.Create(&singer).Error

	return singer, err
}

func (r *repository) UpdateSinger(singer models.Singer) (models.Singer, error) {
	err := r.db.Save(&singer).Error

	return singer, err
}

func (r *repository) DeleteSinger(singer models.Singer) (models.Singer, error) {
	err := r.db.Delete(&singer).Error

	return singer, err
}
