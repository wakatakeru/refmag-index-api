package database

import (
	"github.com/wakatakeru/refmag-index-api/domain"
)

type PaperRepository struct {
	SqlHandler
}

func NewPaperRepository(sqlHandler SqlHandler) PaperRepository {
	paperRepository := PaperRepository{SqlHandler: sqlHandler}
	return paperRepository
}

func (repo *PaperRepository) Store(p domain.Paper) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO papers (title, doi, supplement) VALUES (?,?,?)", p.Title, p.DOI, p.Supplement,
	)

	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err == nil {
		return
	}
	id = int(id64)
	return
}

func (repo *PaperRepository) FindByID(paper_id int) (paper domain.Paper, err error) {
	row, err := repo.Query("SELECT id, title, doi, supplement FROM papers WHERE id = ?", paper_id)
	defer row.Close()
	if err != nil {
		return
	}

	var id int
	var title string
	var doi string
	var supplement string

	row.Next()
	err = row.Scan(&id, &title, &doi, &supplement)
	if err != nil {
		return
	}

	paper.ID = id
	paper.Title = title
	paper.DOI = doi
	paper.Supplement = supplement
	return
}

func (repo *PaperRepository) FindAll() (papers domain.Papers, err error) {
	rows, err := repo.Query("SELECT id, title, doi, supplement FROM papers")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var id int
		var title string
		var doi string
		var supplement string
		if err := rows.Scan(&id, &title, &doi, &supplement); err != nil {
			continue
		}
		paper := domain.Paper{
			ID:         id,
			Title:      title,
			DOI:        doi,
			Supplement: supplement,
		}
		papers = append(papers, paper)
	}
	return
}
