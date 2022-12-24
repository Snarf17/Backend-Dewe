package authdto

type RegisterResponse struct {
	Email string `json:"email" form:"email" validate:"required"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}
type LoginResponse struct {
	Email string `json:"email" form:"email" validate:"required"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}
