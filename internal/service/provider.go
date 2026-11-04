package service

import (
	"context"
	"fmt"

	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
)

type ProviderService struct {
	repo *db.Queries
}

func NewProviderService(repo *db.Queries) *ProviderService {
	return &ProviderService{repo: repo}
}

func (s *ProviderService) GetProviderList(ctx context.Context, params db.GetAllProvidersParams) (*[]db.GetAllProvidersRow, error) {

	fmt.Printf("params: %+v", params)
	providers, err := s.repo.GetAllProviders(ctx, params)
	if err != nil {
		return nil, err
	}
	return &providers, nil

}

func (s *ProviderService) AddCategory(ctx context.Context, data db.AddCategoryParams) (*db.Category, error) {
	category, err := s.repo.AddCategory(ctx, data)
	return &category, err
}

func (s *ProviderService) GetCategories(ctx context.Context) (*[]db.Category, error) {
	categories, err := s.repo.GetAllCategories(ctx)
	return &categories, err
}

func (s *ProviderService) UpdateProvier(ctx context.Context, data db.UpdateProviderParams) (int64, error) {
	n, err := s.repo.UpdateProvider(ctx, data)
	return n, err
}
