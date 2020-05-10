package database

import (
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/wakatakeru/refmag-index-api/domain"
)

func TestStore(t *testing.T) {
	ctrl := gomock.NewController(t)

	paper := domain.Paper{}
	query := "INSERT INTO papers (title, doi, supplement) VALUES (?,?,?)"
	var expectedID int64
	var err error

	mockSqlHandler := NewMockSqlHandler(ctrl)
	mockSqlResult := NewMockSqlResult(ctrl)

	mockSqlHandler.EXPECT().Execute(query, paper.Title, paper.DOI, paper.Supplement).Return(mockSqlResult, err)
	mockSqlResult.EXPECT().LastInsertId().Return(expectedID, err)

	paperRepository := NewPaperRepository(mockSqlHandler)
	_, err = paperRepository.Store(paper)

	if err != nil {
		t.Error("Store is not same as expected")
	}
}

func TestFindByID(t *testing.T) {
	return
}

func TestFindAll(t *testing.T) {
	return
}
