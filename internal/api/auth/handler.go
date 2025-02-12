package auth

import (
	"context"
	std_errors "errors"

	"github.com/AlekSi/pointer"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/generated/api"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/errors"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/logger"
)

type AuthHandler struct {
	log     logger.Log
	usecase authUsecase
}

func New(log logger.Log, usecase authUsecase) *AuthHandler {
	return &AuthHandler{
		log:     log,
		usecase: usecase,
	}
}

// TODO: решить вопрос с ограничением в req
func (h *AuthHandler) APIAuthPost(ctx context.Context, req *api.AuthRequest) (api.APIAuthPostRes, error) {
	token, err := h.usecase.Auth(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		if std_errors.Is(err, errors.ErrInvalidPassword) {
			return pointer.To(api.APIAuthPostUnauthorized(api.ErrorResponse{Errors: api.NewOptString(err.Error())})), nil
		}
		h.log.WithContext(ctx).WithError(err).Error(api.APIAuthPostOperation)
		return pointer.To(api.APIAuthPostInternalServerError(api.ErrorResponse{Errors: api.NewOptString(errors.ErrInternal.Error())})), nil
	}

	return &api.AuthResponse{Token: api.NewOptString(token)}, nil
}
