package domain

import (
	"time"
)

type User struct {
    ID        string    `json:"id"`
    Email     string    `json:"email"`
    Name      string    `json:"name"`
    Role      string    `json:"role"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    LastLogin time.Time `json:"last_login"`
}

type UserRegister struct {
		Email     string    `json:"email" binding:"required"`
    Name      string    `json:"name" binding:"required"`
		Phone 		string 		`json:"phone" binding:"required"`
		Password	string 		`json:"password" binding:"required"`
}
