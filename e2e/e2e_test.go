package e2e

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/require"
)

const baseURL = "http://localhost:8080"

func authenticateUser(e *httpexpect.Expect, username, password string) string {
	resp := e.POST("/api/auth").
		WithJSON(map[string]string{
			"username": username,
			"password": password,
		}).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	return resp.Value("token").String().Raw()
}

func TestE2E_SendCoins(t *testing.T) {
	e := httpexpect.New(t, baseURL)

	token1 := authenticateUser(e, "user1", "password123")
	authenticateUser(e, "user2", "password123")

	infoBefore := e.GET("/api/info").
		WithHeader("Authorization", "Bearer "+token1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	coinsBefore := infoBefore.Value("coins").Number().Raw()

	e.POST("/api/sendCoin").
		WithHeader("Authorization", "Bearer "+token1).
		WithJSON(map[string]interface{}{
			"toUser": "user2",
			"amount": 10,
		}).
		Expect().
		Status(http.StatusOK)

	infoAfter := e.GET("/api/info").
		WithHeader("Authorization", "Bearer "+token1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	coinsAfter := infoAfter.Value("coins").Number().Raw()
	require.Equal(t, coinsBefore-10, coinsAfter)
}

func TestE2E_BuyMerch(t *testing.T) {
	e := httpexpect.New(t, baseURL)

	token1 := authenticateUser(e, "user1", "password123")

	e.GET("/api/buy/cup").
		WithHeader("Authorization", "Bearer "+token1).
		Expect().
		Status(http.StatusOK)

	infoAfterBuy := e.GET("/api/info").
		WithHeader("Authorization", "Bearer "+token1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	inventory := infoAfterBuy.Value("inventory").Array()
	require.NotEmpty(t, inventory)
}
