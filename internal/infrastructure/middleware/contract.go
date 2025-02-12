package middleware

import "github.com/google/uuid"

type jwtToken interface {
	GetUserIDFromToken(token string) (uuid.UUID, error)
}
