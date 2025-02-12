// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// APIAuthPost implements POST /api/auth operation.
	//
	// Аутентификация и получение JWT-токена. При первой
	// аутентификации пользователь создается
	// автоматически.
	//
	// POST /api/auth
	APIAuthPost(ctx context.Context, req *AuthRequest) (APIAuthPostRes, error)
	// APIBuyItemGet implements GET /api/buy/{item} operation.
	//
	// Купить предмет за монеты.
	//
	// GET /api/buy/{item}
	APIBuyItemGet(ctx context.Context, params APIBuyItemGetParams) (APIBuyItemGetRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
