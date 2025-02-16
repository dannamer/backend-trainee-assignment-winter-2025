package buyitem

import (
	"context"
	std_errors "errors"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/generated/api"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/errors"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/logger"
)

type BuyItemHandler struct {
	log     logger.Log
	usecase buyItemUsecase
}

func New(log logger.Log, usecase buyItemUsecase) *BuyItemHandler {
	return &BuyItemHandler{
		log:     log,
		usecase: usecase,
	}
}

func (h *BuyItemHandler) APIBuyItemGet(ctx context.Context, params api.APIBuyItemGetParams) (api.APIBuyItemGetRes, error) {
	userID, _ := ctx.Value(domain.UserIDKey).(uuid.UUID)
	if err := h.usecase.BuyItem(ctx, userID, params.Item); err != nil {
		if std_errors.Is(err, errors.ErrInsufficientFound) {
			return pointer.To(api.APIBuyItemGetBadRequest(api.ErrorResponse{Errors: api.NewOptString(errors.ErrInsufficientFound.Error())})), nil
		}
		if std_errors.Is(err, errors.ErrMerchNotFound) {
			return pointer.To(api.APIBuyItemGetBadRequest(api.ErrorResponse{Errors: api.NewOptString(errors.ErrMerchNotFound.Error())})), nil
		}
		h.log.WithContext(ctx).WithError(err).Error(api.APIBuyItemGetOperation)
		return pointer.To(api.APIBuyItemGetInternalServerError(api.ErrorResponse{Errors: api.NewOptString(errors.ErrInternal.Error())})), nil
	}

	return &api.APIBuyItemGetOK{}, nil
}
