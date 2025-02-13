package info_storage

import (
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/postgres"
)

type storage struct {
	pg postgres.PgxPool
}

func New(pg postgres.PgxPool) *storage {
	return &storage{
		pg: pg,
	}
}
