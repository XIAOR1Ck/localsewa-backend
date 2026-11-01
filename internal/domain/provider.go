package domain

import (
	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
)

// func DbProviderToModel(u db.Provider) Provider {
// 	return Provider{
// 		ID:        u.ID.String(),
// 		Email:     u.Email,
// 		Name:      u.Name,
// 		Phone:     u.Phone,
// 		CreatedAt: u.CreatedAt.Time,
// 		UpdatedAt: u.UpdatedAt.Time,
// 	}
// }

// type Provider struct {
// 	ID           string    `json:"id"`
// 	Name         string    `json:"name"`
// 	Email        string    `json:"email"`
// 	Phone        string    `json:"phone"`
// 	ProfileImage string    `json:"profileImage"`
// 	Category     string    `json:"category"`
// 	Experience   float32   `json:"experience"`
// 	HourlyRate   float32   `json:"hourlyRate"`
// 	Description  string    `json:"description"`
// 	Rating       float32   `json:"rating"`
// 	TotalReviews int       `json:"totalReviews"`
// 	CreatedAt    time.Time `json:"created_at"`
// 	UpdatedAt    time.Time `json:"updated_at"`
// }

type ProviderRegisterModel struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProviderBuilder struct {
	provider *db.Provider
}

func (b *ProviderBuilder) SetName(name string) *ProviderBuilder {
	b.provider.Name = name
	return b
}
func (b *ProviderBuilder) SetCategory(id string) *ProviderBuilder {
	b.provider.Category = &id
	return b
}
func (b *ProviderBuilder) SetHourlyRate(rate float32) *ProviderBuilder {
	b.provider.HourlyRate = &rate
	return b
}
func (b *ProviderBuilder) SetEmail(email string) *ProviderBuilder {
	b.provider.Email = email
	return b
}

func (b *ProviderBuilder) SetDescription(desc string) *ProviderBuilder {
	b.provider.Description = &desc
	return b
}

func (b *ProviderBuilder) SetExperience(experience float32) *ProviderBuilder {
	b.provider.Experience = experience
	return b
}

func (b *ProviderBuilder) SetPhone(phone string) *ProviderBuilder {
	b.provider.Phone = phone
	return b
}

func (b *ProviderBuilder) Build() *db.Provider {
	return b.provider
}
