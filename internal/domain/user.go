package domain

import (
	"time"

	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
)


func DbUserToModel(u db.User) User {
	return User{
		ID: u.ID.String(),
		Email: u.Email,
		Name: u.Name,
		Role: string(u.Role),
		Phone: u.Phone,
		CreatedAt: u.CreatedAt.Time,
		UpdatedAt: u.UpdatedAt.Time,
	}
}


type User struct {
    ID        string    `json:"id"`
    Email     string    `json:"email"`
    Name      string    `json:"name"`
		Phone 		string 		`json:"phone"`
    Role      string    `json:"role"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type UserRegisterModel struct {
		Email     string    `json:"email" binding:"required"`
    Name      string    `json:"name" binding:"required"`
		Phone 		string 		`json:"phone" binding:"required"`
		Password	string 		`json:"password" binding:"required"`
}

type UpdateUserModel struct {
	Email 			*string 		`json:"email"`
	Name 				*string 		`json:"name"`
	Role 				*string   	`json:"role"`
}

func (u User) UpdateWith(opts UpdateUserModel) User {
	if opts.Email != nil {
		u.Email = *opts.Email
	}
	if opts.Name != nil {
		u.Name = *opts.Name
	}
	if opts.Role != nil {
		u.Role = *opts.Role
	}

	return u
}
