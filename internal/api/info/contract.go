package info

import (
	"context"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
)


type InfoUsecase interface {
	GetInfo(ctx context.Context) (domain.Info, error)
}