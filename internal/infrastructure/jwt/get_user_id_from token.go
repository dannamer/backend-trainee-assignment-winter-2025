package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (j *JwtToken) GetUserIDFromToken(token string) (uuid.UUID, error) {
	parsedToken, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(j.JwtKey), nil
		},
	)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return uuid.UUID{}, fmt.Errorf("invalid token")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("id not found in token")
	}

	id, err := uuid.Parse(sub)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("invalid ID format in token: %w", err)
	}
	return id, nil
}
