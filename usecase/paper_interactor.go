package usecase

import "github.com/wakatakeru/refmag-index-api/domain"

type PaperInteractor struct {
	PaperRepository PaperRepository
}

func NewPaperInteractor(paperRepository PaperRepository) PaperInteractor {
	paperInteractor := PaperInteractor{PaperRepository: paperRepository}
	return paperInteractor
}

func (interactor *PaperInteractor) Add(p domain.Paper) (err error) {
	_, err = interactor.PaperRepository.Store(p)
	return
}

func (interactor *PaperInteractor) Paper(id int) (paper domain.Paper, err error) {
	paper, err = interactor.PaperRepository.FindByID(id)
	return
}

func (interactor *PaperInteractor) Papers() (papers domain.Papers, err error) {
	papers, err = interactor.PaperRepository.FindAll()
	return
}
