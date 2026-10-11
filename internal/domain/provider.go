package domain

import (
	"time"
	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
)


func DbProviderToModel(u db.Provider) Provider {
	return Provider{
		ID: u.ID.String(),
		Email: u.Email,
		Name: u.Name,
		Phone: u.Phone,
		CreatedAt: u.CreatedAt.Time,
		UpdatedAt: u.UpdatedAt.Time,
	}
}

type Provider struct {
	ID 						string 				`json:"id"`
	Name					string 				`json:"name"`
	Email   			string 				`json:"email"`
	Phone  				string 				`json:"phone"`
	ProfileImage 	*string 			`json:"profileImage"`
	Category 			*string 			`json:"category"`
	Experience 		*float32 			`json:"experience"`
	HourlyRate		*float32 			`json:"hourlyRate"`	
	Description 	*string 			`json:"description"`
	Rating 				*float32 			`json:"rating"`
	TotalReviews 	*int 					`json:"totalReviews"`
	CreatedAt 		time.Time 		`json:"created_at"`
  UpdatedAt 		time.Time 		`json:"updated_at"`

}


type ProviderRegisterModel struct {
		Email     string    `json:"email" binding:"required"`
    Name      string    `json:"name" binding:"required"`
		Phone 		string 		`json:"phone" binding:"required"`
		Password	string 		`json:"password" binding:"required"`
}

type UpdateProviderModel struct {
	Email 			*string 		`json:"email"`
	Name 				*string 		`json:"name"`
	Category		*string 		`json:"category"`
	Description *string 		`json:"description"`
	Phone 			*string 		`json:"phone"`
	HourlyRate	*float32 	`json:"hourlyRate"`	

}


