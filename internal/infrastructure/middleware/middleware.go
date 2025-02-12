package middleware

import (
	"context"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/generated/api"
)

type Middleware struct {
	jwt jwtToken
}

func New(jwt jwtToken) *Middleware {
	return &Middleware{
		jwt: jwt,
	}
}

func (h *Middleware) HandleBearerAuth(ctx context.Context, operationName api.OperationName, t api.BearerAuth) (context.Context, error) {
	id, err := h.jwt.GetUserIDFromToken(t.GetToken())
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, domain.UserIDKey, id)

	return ctx, nil
}