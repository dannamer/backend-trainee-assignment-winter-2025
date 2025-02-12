package jwt

type JwtToken struct {
	JwtKey string
}

func New(key string) *JwtToken {
	return &JwtToken{
		JwtKey: key,
	}
}
