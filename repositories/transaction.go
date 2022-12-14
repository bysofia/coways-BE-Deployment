package repositories

import (
	"BackEnd/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	FindTransactions() ([]models.Transaction, error)
	GetTransactionById(ID int) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
	GetUserId(UserID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Create(&transaction).Error

	return transaction, err
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionById(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) GetUserId(UserID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Where("user_id = ? AND subscription = ?", UserID, "Active").First(&transaction).Error
	return transaction, err
}

func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.First(&transaction, ID)

	if status != transaction.Status && status == "success" {
		transaction.Status = status
		transaction.Remaining = 30
		transaction.Subscription = "Active"
	}

	if status != transaction.Status && status == "failed" {
		transaction.Status = status
		transaction.Remaining = 0
		transaction.Subscription = "Not Active"
	}

	err := r.db.Save(&transaction).Error
	return err
}
