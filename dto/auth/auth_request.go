package authdto

type ResgisterRequest struct {
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    int    `json:"phone" gorm:"type: int(100)" validate:"required"`
	Address  string `json:"address" gorm:"type: text" validate:"required"`
}
type LoginRequest struct {
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    int    `json:"phone" gorm:"type: int(100)" validate:"required"`
	Address  string `json:"address" gorm:"type: text" validate:"required"`
}
