package info

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/generated/api"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/errors"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/logger"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type InfoHandler struct {
	log     logger.Log
	usecase InfoUsecase
}

func New(log logger.Log, usecase InfoUsecase) *InfoHandler {
	return &InfoHandler{
		log:     log,
		usecase: usecase,
	}
}

func (h *InfoHandler) APIInfoGet(ctx context.Context) (api.APIInfoGetRes, error) {
	userID, _ := ctx.Value(domain.UserIDKey).(uuid.UUID)
	info, err := h.usecase.GetInfo(ctx, userID)
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error(api.APIInfoGetOperation)
		return pointer.To(api.APIInfoGetInternalServerError(api.APIInfoGetInternalServerError{Errors: api.NewOptString(errors.ErrInternal.Error())})), nil
	}

	response := api.InfoResponse{
		Coins: api.NewOptInt(info.Coin),
		Inventory: lo.Map(info.Inventory, func(item domain.Inventory, _ int) api.InfoResponseInventoryItem {
			return api.InfoResponseInventoryItem{
				Type: api.OptString{
					Value: item.Item,
					Set:   item.Item != "",
				},
				Quantity: api.OptInt{
					Value: item.Quantity,
					Set:   item.Quantity != 0,
				},
			}
		}),
		CoinHistory: api.OptInfoResponseCoinHistory{
			Value: api.InfoResponseCoinHistory{
				Received: lo.Map(info.CoinHistory.Received, func(received domain.Transaction, _ int) api.InfoResponseCoinHistoryReceivedItem {
					return api.InfoResponseCoinHistoryReceivedItem{
						FromUser: api.OptString{
							Value: received.Username,
							Set:   received.Username != "",
						},
						Amount: api.OptInt{
							Value: received.Amount,
							Set:   received.Username != "",
						},
					}
				}),
				Sent: lo.Map(info.CoinHistory.Sent, func(sent domain.Transaction, _ int) api.InfoResponseCoinHistorySentItem {
					return api.InfoResponseCoinHistorySentItem{
						ToUser: api.OptString{
							Value: sent.Username,
							Set:   sent.Username != "",
						},
						Amount: api.OptInt{
							Value: sent.Amount,
							Set:   sent.Username != "",
						},
					}
				}),
			},
			Set: len(info.CoinHistory.Received) != 0 || len(info.CoinHistory.Sent) != 0,
		},
	}
	return &response, nil
}
