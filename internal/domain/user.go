package domain

import (
	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
)

type UserRegisterModel struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserModel struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
	Role  *string `json:"role"`
}

type UserBuilder struct {
	user *db.User
}

// Builder function
func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: &db.User{},
	}
}

// Builders
func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) WithPhone(phone string) *UserBuilder {
	b.user.Phone = phone
	return b
}

func (b *UserBuilder) WithProfileImage(img string) *UserBuilder {
	b.user.ProfileImage = &img
	return b
}

func (b *UserBuilder) Build() *db.User {
	return b.user
}
