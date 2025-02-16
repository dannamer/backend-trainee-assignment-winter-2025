package sendcoin_storage

import (
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/postgres"
)

type storage struct {
	pg       postgres.PgxPool
	trGetter *trmpgx.CtxGetter
}

func New(pg postgres.PgxPool) *storage {
	return &storage{
		pg:       pg,
		trGetter: trmpgx.DefaultCtxGetter,
	}
}
