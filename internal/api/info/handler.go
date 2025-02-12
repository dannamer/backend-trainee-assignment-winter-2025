package info

// import (
// 	"context"

// 	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
// 	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/generated/api"
// 	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/logger"
// 	"github.com/samber/lo"
// )

// type InfoHandler struct {
// 	log     logger.Log
// 	usecase InfoUsecase
// }

// func New(log logger.Log, usecase InfoUsecase) *InfoHandler {
// 	return &InfoHandler{
// 		log:     log,
// 		usecase: usecase,
// 	}
// }

// func (h *InfoHandler) APIInfoGet(ctx context.Context) (api.APIInfoGetRes, error) {
// 	info, err := h.usecase.GetInfo(ctx)
// 	if err != nil {

// 	}
// 	// response := lo.Map(devs, func(dev development.Development, _ int) api.DevelopmentSearchBoardOKItem {
// 	// 	return api.DevelopmentSearchBoardOKItem{
// 	// 		ID:   dev.ID,
// 	// 		Name: dev.Name,
// 	// 		Coords: api.DevelopmentSearchBoardOKItemCoords{
// 	// 			Lat: dev.Coordinates.Lat(),
// 	// 			Lon: dev.Coordinates.Lon(),
// 	// 		},
// 	// 		ImageUrl:    lo.If(dev.Meta.ImageURL != "", dev.Meta.ImageURL).Else(defaultImageUrl),
// 	// 		Description: dev.Meta.Description,
// 	// 	}
// 	// })

// 	response := api.InfoResponse{
// 		Coins: api.OptInt{
// 			Value: int(info.Coin.IntPart()),
// 			Set:   true,
// 		},
// 		Inventory: lo.Map(info.Inventory, func(item domain.Inventory, _ int) api.InfoResponseInventoryItem {
// 			return api.InfoResponseInventoryItem{
// 				Type: api.OptString{
// 					Value: item.Item,
// 					Set:   item.Item != "",
// 				},
// 				Quantity: api.OptInt{
// 					Value: item.Quantity,
// 					Set:   item.Quantity != 0,
// 				},
// 			}
// 		}),
// 		CoinHistory: api.OptInfoResponseCoinHistory{
// 			Value: api.InfoResponseCoinHistory{
// 				Received: lo.Map(info.CoinHistory.Received, func(received domain.Received, _ int) api.InfoResponseCoinHistoryReceivedItem {
// 					return api.InfoResponseCoinHistoryReceivedItem{
// 						FromUser: api.OptString{
// 							Value: received.FromUser,
// 							Set:   received.FromUser != "",
// 						},
// 						Amount: api.OptInt{
// 							Value: int(received.Amount.IntPart()),
// 							Set:   received.FromUser != "",
// 						},
// 					}
// 				}),
// 			},
// 			Set: true,
// 		},
// 	}
// 	return &response, nil
// }
