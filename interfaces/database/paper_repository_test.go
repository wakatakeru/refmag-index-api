package database

import (
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/wakatakeru/refmag-index-api/domain"
)

func TestStore(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockSqlHandler := NewMockSqlHandler(ctrl)
	mockSqlResult := NewMockSqlResult(ctrl)

	paper := domain.Paper{}
	query := "INSERT INTO papers (title, doi, supplement) VALUES (?,?,?)"
	var expectedID int64
	var err error

	mockSqlHandler.EXPECT().Execute(query, paper.Title, paper.DOI, paper.Supplement).Return(mockSqlResult, err)
	mockSqlResult.EXPECT().LastInsertId().Return(expectedID, err)

	paperRepository := NewPaperRepository(mockSqlHandler)
	_, err = paperRepository.Store(paper)

	if err != nil {
		t.Error("Store is not same as expected")
	}
}

func TestFindByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockSqlHandler := NewMockSqlHandler(ctrl)
	mockRow := NewMockRow(ctrl)

	query := "SELECT id, title, doi, supplement FROM papers WHERE id = ?"
	var paperID int
	var err error
	var id int
	var title string
	var doi string
	var supplement string

	mockSqlHandler.EXPECT().Query(query, paperID).Return(mockRow, err)
	mockRow.EXPECT().Next()
	mockRow.EXPECT().Scan(&id, &title, &doi, &supplement).Return(err)
	mockRow.EXPECT().Close()

	paperRepository := NewPaperRepository(mockSqlHandler)
	_, err = paperRepository.FindByID(paperID)

	if err != nil {
		t.Error("FindByID is not same as expected")
	}
}

func TestFindAll(t *testing.T) {
	return
}
