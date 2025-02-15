package info_usecase

type InfoUsecase struct {
	storage   infoStorage
}

func New(storage infoStorage) *InfoUsecase {
	return &InfoUsecase{
		storage:   storage,
	}
}
