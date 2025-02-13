package buyitem_usecase

type buyItemUsecase struct {
	storage   buyItemStorage
	trManager trManager
}

func New(storage buyItemStorage, trManager trManager) *buyItemUsecase {
	return &buyItemUsecase{
		storage:   storage,
		trManager: trManager,
	}
}
