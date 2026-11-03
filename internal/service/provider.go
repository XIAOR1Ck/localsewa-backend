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

func (s *ProviderService) GetProviderList(ctx context.Context, limit int32, offset int32) (*[]db.GetAllProviderRow, error) {
	params := db.GetAllProviderParams{
		Limit:  limit,
		Offset: offset,
	}
	fmt.Printf("params: %+v", params)
	providers, err := s.repo.GetAllProvider(ctx, params)
	if err != nil {
		return nil, err
	}
	return &providers, nil

}
