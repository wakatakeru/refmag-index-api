package usecase

import "github.com/wakatakeru/refmag-index-api/domain"

type PaperInteractor struct {
	PaperRepository PaperRepository
}

func (interactor *PaperInteractor) Add(p domain.Paper) (err error) {
	_, err = interactor.PaperRepository.Store(p)
	return
}

func (interactor *PaperInteractor) Papers() (paper domain.Papers, err error) {
	paper, err = interactor.PaperRepository.FindAll()
	return
}
