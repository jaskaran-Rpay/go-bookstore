package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	First_Name    string    `json:"first_Name" validate:"required, min = 2, max = 100"`
	Last_Name     string    `json:"last_Name" validate:"required, min = 2, max = 100"`
	Email         string    `json:"email" validate:"email, required"`
	Password      string    `json:"password" validate:"required, min = 8"`
	Phone         string    `json:"phone" validate:"required"`
	User_type     string    `json:"user_type" validate:"required, eq = ADMIN | eq = USER"`
	Token         string    `json:"token"`
	Refresh_Token string    `json:"refresh_token"`
	Created_At    time.Time `json:"created_at"`
	Updated_At    time.Time `json:"updated_at"`
}
