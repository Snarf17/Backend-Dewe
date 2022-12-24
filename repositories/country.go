package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type CountryRepository interface {
	FindCountry() ([]models.Country, error)
	GetCountry(ID int) (models.Country, error)
	CreateCountry(country models.Country) (models.Country, error)
	UpdateCountry(country models.Country) (models.Country, error)
	// DeleteCountry(user models.User, ID int) (models.User, error)
}

type repositoryCountry struct {
	db *gorm.DB
}

func RepositoryCountry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCountry() ([]models.Country, error) {
	var country []models.Country
	err := r.db.Find(&country).Error

	return country, err
}
func (r *repository) GetCountry(ID int) (models.Country, error) {
	var country models.Country
	err := r.db.First(&country, ID).Error

	return country, err
}
func (r *repository) CreateCountry(country models.Country) (models.Country, error) {
	// var user models.User
	err := r.db.Create(&country).Error

	return country, err
}
func (r *repository) UpdateCountry(country models.Country) (models.Country, error) {
	// var Country models.Country
	err := r.db.Save(&country).Error

	return country, err
}
