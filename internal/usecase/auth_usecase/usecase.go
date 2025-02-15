package auth_usecase

type AuthUsecase struct {
	storage   authStorage
	trManager trManager
	jwt       jwtToken
	password  password
}

func New(storage authStorage, trManager trManager, jwt jwtToken, password password) *AuthUsecase {
	return &AuthUsecase{
		storage:   storage,
		trManager: trManager,
		jwt:       jwt,
		password: password,
	}
}
