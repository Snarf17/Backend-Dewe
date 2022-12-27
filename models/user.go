package models

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Phone    int    `json:"phone" gorm:"type: int(100)"`
	Address  string `json:"address" gorm:"type: text"`
	Role     string `json:"role"`
}

// type UserResponse struct {
// 	ID       int    `json:"id"`
// 	FullName string `json:"fullname"`
// 	Email    string `json:"email"`
// 	// Password string `json:"password"`
// 	Phone   int    `json:"phone"`
// 	Address string `json:"address"`
// }

// func (UserResponse) TableName() string {
// 	return "users"
// }
