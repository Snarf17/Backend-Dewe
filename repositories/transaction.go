package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	// GetUsers(ID int) (models.User, error)
	// CreateUsers(user models.User) (models.User, error)
	// UpdateUsers(user models.User, ID int) (models.User, error)
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
	err := r.db.Find(&transaction).Error

	return transaction, err
}
