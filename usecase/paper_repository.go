package usecase

import "github.com/wakatakeru/refmag-index-api/domain"

type PaperRepository interface {
	Store(domain.Paper) (int, error)
	FindAll() (domain.Papers, error)
}
