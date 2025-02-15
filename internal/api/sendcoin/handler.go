package sendcoin

import (
	"context"
	std_errors "errors"

	"github.com/AlekSi/pointer"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/generated/api"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/errors"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/logger"
	"github.com/google/uuid"
)

type SendCoinHandler struct {
	log     logger.Log
	usecase sendcoinUsecase
}

func New(log logger.Log, usecase sendcoinUsecase) *SendCoinHandler {
	return &SendCoinHandler{
		log:     log,
		usecase: usecase,
	}
}

func (h *SendCoinHandler) APISendCoinPost(ctx context.Context, req *api.SendCoinRequest) (api.APISendCoinPostRes, error) {
	userID, _ := ctx.Value(domain.UserIDKey).(uuid.UUID)
	if err := h.usecase.SendCoin(ctx, req.GetToUser(), userID, int64(req.GetAmount())); err != nil {
		if std_errors.Is(err, errors.ErrUserNotFound) {
			return pointer.To(api.APISendCoinPostBadRequest(api.ErrorResponse{Errors: api.NewOptString(errors.ErrUserNotFound.Error())})), nil
		}
		if std_errors.Is(err, errors.ErrInsufficientFound) {
			return pointer.To(api.APISendCoinPostBadRequest(api.ErrorResponse{Errors: api.NewOptString(errors.ErrInsufficientFound.Error())})), nil
		}
		if std_errors.Is(err, errors.ErrSelfTransfer) {
			return pointer.To(api.APISendCoinPostBadRequest(api.ErrorResponse{Errors: api.NewOptString(errors.ErrSelfTransfer.Error())})), nil
		}
		h.log.WithContext(ctx).WithError(err).Error(api.APISendCoinPostOperation)
		return pointer.To(api.APISendCoinPostInternalServerError(api.ErrorResponse{Errors: api.NewOptString(errors.ErrInternal.Error())})), nil
	}

	return &api.APISendCoinPostOK{}, nil
}
