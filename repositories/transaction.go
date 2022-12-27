package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	// DeleteUsers(user models.User, ID int) (models.User, error)
}

type repositoryTransaction struct {
	db *gorm.DB
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var Transaction models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").First(&Transaction, ID).Error

	return Transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	// var user models.User
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(Transaction models.Transaction) (models.Transaction, error) {
	// var user models.User
	err := r.db.Model(&Transaction).Updates(Transaction).Error

	return Transaction, err
}

// func (r *repository) DeleteTransaction(Transaction models.Transaction, ID int) (models.Transaction, error) {
// 	// var user models.User
// 	err := r.db.Preload("Country").Delete(&Transaction).Error

// 	return Transaction, err
// }
