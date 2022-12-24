package tripdto

import "time"

type TripResponse struct {
	// ID       int    `json:"id"`
	Title          string    `json:"title" form:"title" validate:"required"`
	CountryID      int       `json:"country_id" form:"country_id" `
	Accomodation   string    `json:"accomodation" form:"accomodation" `
	Transportation string    `json:"transport" form:"transport" `
	Eat            string    `json:"eat" form:"eat" `
	Day            int       `json:"day" form:"day" `
	Night          int       `json:"night" form:"night" `
	DateTrip       time.Time `json:"date_trip" form:"date_trip" `
	Price          int       `json:"price" form:"price" `
	Quota          int       `json:"quota" form:"quota" `
	Description    string    `json:"desc" form:"desc" `
	Image          string    `json:"image" form:"image" `
}

type TripDeleteResponse struct {
	ID int `json:"id"`
}
