package models

type Country struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment"`
	Name   string `json:"name"`
	UserID int    `json:"-"`
}

type CountryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (CountryResponse) TableName() string {
	return "countries"
}
