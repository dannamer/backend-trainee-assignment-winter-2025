package sendcoin_usecase

type sendCoinUsecase struct {
	storage   sendCoinStorage
	trManager trManager
}

func New(storage sendCoinStorage, trManager trManager) *sendCoinUsecase {
	return &sendCoinUsecase{
		storage:   storage,
		trManager: trManager,
	}
}
