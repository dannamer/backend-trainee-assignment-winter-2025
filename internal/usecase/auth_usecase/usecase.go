package auth_usecase

type authUsecase struct {
	storage   authStorage
	trManager trManager
	jwt       jwtToken
}

func New(storage authStorage, trManager trManager, jwt jwtToken) *authUsecase {
	return &authUsecase{
		storage:   storage,
		trManager: trManager,
		jwt:       jwt,
	}
}
