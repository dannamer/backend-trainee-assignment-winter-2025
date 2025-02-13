package info

import (
	"context"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/google/uuid"
)


type InfoUsecase interface {
	GetInfo(ctx context.Context, userID uuid.UUID) (domain.Info, error)
}