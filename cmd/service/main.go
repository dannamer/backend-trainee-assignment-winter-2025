package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"

	trmpgxs "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	serviceapi "github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api/auth"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api/buyitem"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api/info"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api/sendcoin"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/config"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/generated/api"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/jwt"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/logger"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/middleware"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/password"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/postgres"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/storage/auth_storage"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/storage/buyitem_storage"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/storage/info_storage"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/storage/sendcoin_storage"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/auth_usecase"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/buyitem_usecase"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/info_usecase"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/sendcoin_usecase"
)

func main() {
	ctx := context.Background()
	log := logger.NewLogger(slog.LevelDebug, "dev", os.Stdout)
	config := config.MustNewConfigWithEnv()
	jwt := jwt.New(config.JwtKey())
	password := password.New()

	pg, err := postgres.New(ctx, config.PgUrl())
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to connect to postgres")
		return
	}
	defer pg.Close()

	trManager := manager.Must(trmpgxs.NewDefaultFactory(pg.Pool))

	authStorage := auth_storage.New(pg.Pool)
	buyItemStorage := buyitem_storage.New(pg.Pool)
	sendCoinStorage := sendcoin_storage.New(pg.Pool)
	infoStorage := info_storage.New(pg.Pool)

	authUsecase := auth_usecase.New(authStorage, trManager, jwt, password)
	buyItemUsecase := buyitem_usecase.New(buyItemStorage, trManager)
	sendCointUsecase := sendcoin_usecase.New(sendCoinStorage, trManager)
	infoUsecase := info_usecase.New(infoStorage)

	srv := &serviceapi.API{
		AuthHandler:     auth.New(log, authUsecase),
		BuyItemHandler:  buyitem.New(log, buyItemUsecase),
		SendCoinHandler: sendcoin.New(log, sendCointUsecase),
		InfoHandler:     info.New(log, infoUsecase),
	}

	middleware := middleware.New(jwt)

	server, err := api.NewServer(srv, middleware)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to create server")
		os.Exit(1)
	}

	log.WithContext(ctx).Info("server start")

	srvAddr := fmt.Sprintf(":%s", config.HttpPort())
	httpServer := &http.Server{
		Addr:    srvAddr,
		Handler: server,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithContext(ctx).WithError(err).Error("failed to start server")
			os.Exit(1)
		}
	}()

	<-quit
	log.WithContext(ctx).Info("Shutdown signal received, starting graceful shutdown...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.WithContext(ctx).WithError(err).Error("Server forced to shutdown")
	} else {
		log.WithContext(ctx).Info("Server gracefully shut down")
	}

}
