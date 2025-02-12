package buyitem_usecase

type BuyItemUsecase struct {
	storage   buyItemStorage
	trManager trManager
}

func New(storage buyItemStorage, trManager trManager) *BuyItemUsecase {
	return &BuyItemUsecase{
		storage:   storage,
		trManager: trManager,
	}
}
