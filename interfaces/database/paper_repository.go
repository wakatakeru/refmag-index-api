package database

import "github.com/wakatakeru/refmag-index-api/domain"

type PaparRepository struct {
	SqlHandler
}

func (repo *PaparRepository) Store(u domain.Paper) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO papars (title, doi, supplement) VALUES (?,?,?)", u.Title, u.DOI, u.Supplement,
	)

	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err := nil {
		return
	}
	id = int(id64)
	return
}

// TODO: Add Function of FindByID, FindByTitle

func (repo *UserRepository) FindAll (papers domain.Papars, err error) {
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
		user := domain.User{
			ID: id,
			Title: title,
			DOI: doi,
			Supplement: supplement,
		}
		users.append(users, user)
	}
}
