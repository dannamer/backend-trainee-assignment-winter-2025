package sendcoin_usecase

type SendCoinUsecase struct {
	storage   sendCoinStorage
	trManager trManager
}

func New(storage sendCoinStorage, trManager trManager) *SendCoinUsecase {
	return &SendCoinUsecase{
		storage:   storage,
		trManager: trManager,
	}
}
