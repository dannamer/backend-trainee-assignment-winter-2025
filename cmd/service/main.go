package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"

	trmpgxs "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	serviceapi "github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api/auth"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api/buyitem"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/config"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/generated/api"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/jwt"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/logger"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/middleware"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/postgres"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/storage/auth_storage"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/storage/buyitem_storage"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/auth_usecase"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/buyitem_usecase"
)

func main() {
	ctx := context.Background()
	log := logger.NewLogger(slog.LevelDebug, "dev", os.Stdout)
	config := config.MustNewConfigWithEnv()
	jwt := jwt.New(config.JwtKey())

	pg, err := postgres.New(config.PgUrl())
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to connect to postgres")
	}
	defer pg.Close()

	trManager := manager.Must(trmpgxs.NewDefaultFactory(pg.Pool))

	authStorage := auth_storage.New(pg.Pool)
	buyitemStorage := buyitem_storage.New(pg.Pool)

	authUsecase := auth_usecase.New(authStorage, trManager, jwt)

	buyitemUsecase := buyitem_usecase.New(buyitemStorage, trManager)

	srv := &serviceapi.API{
		AuthHandler: auth.New(log, authUsecase),
		BuyItemHandler: buyitem.New(log, buyitemUsecase),
	}
	
	middleware := middleware.New(jwt)

	server, err := api.NewServer(srv, middleware)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to create server")
		os.Exit(1)
	}

	log.WithContext(ctx).Info("server start")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.HttpPort()), server); err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to start server")
		os.Exit(1)
	}
}
