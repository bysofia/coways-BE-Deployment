package repositories

import (
	"BackEnd/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var user []models.User
	err := r.db.Debug().Preload("Profile").Find(&user).Error

	return user, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Debug().Preload("Profile").First(&user, ID).Error

	return user, err

}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Debug().Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Debug().Delete(&user).Error

	return user, err
}
