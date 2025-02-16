package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (j *JwtToken) GenerateJWT(ID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{}
	claims["sub"] = ID.String()
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.JwtKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
