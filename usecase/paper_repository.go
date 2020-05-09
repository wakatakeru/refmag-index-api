package usecase

import "github.com/wakatakeru/refmag-index-api/domain"

type PaperRepository interface {
	Store(domain.Paper) (int, error)
	FindByID(int) (domain.Paper, error)
	FindAll() (domain.Papers, error)
}
