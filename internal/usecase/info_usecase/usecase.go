package info_usecase

type infoUsecase struct {
	storage   infoStorage
}

func New(storage infoStorage) *infoUsecase {
	return &infoUsecase{
		storage:   storage,
	}
}
