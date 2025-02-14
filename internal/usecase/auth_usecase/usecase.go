package auth_usecase

type authUsecase struct {
	storage   authStorage
	trManager trManager
	jwt       jwtToken
	password  password
}

func New(storage authStorage, trManager trManager, jwt jwtToken, password password) *authUsecase {
	return &authUsecase{
		storage:   storage,
		trManager: trManager,
		jwt:       jwt,
		password: password,
	}
}
