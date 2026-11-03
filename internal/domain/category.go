package domain

import (
	"math/rand"
)

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func generateRandomID() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 12)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewCategory(name string, description string) *Category {
	return &Category{
		ID:          generateRandomID(),
		Name:        name,
		Description: description,
	}
}
