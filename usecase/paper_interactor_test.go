package usecase

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/wakatakeru/refmag-index-api/domain"
	mock "github.com/wakatakeru/refmag-index-api/mock"
)

func TestPaper(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected domain.Paper
	var err error
	id := 0

	mockSample := mock.NewMockPaperRepository(ctrl)
	mockSample.EXPECT().FindByID(id).Return(expected, err)

	paperInteractor := NewPaperInteractor(mockSample)
	result, err := paperInteractor.Paper(id)

	if err != nil {
		t.Error("FindByID is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Error("FindByID is not same as expected")
	}
}

func TestPapers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected domain.Papers
	var err error

	mockSample := mock.NewMockPaperRepository(ctrl)
	mockSample.EXPECT().FindAll().Return(expected, err)

	paperInteractor := NewPaperInteractor(mockSample)
	result, err := paperInteractor.Papers()

	if err != nil {
		t.Error("FindAll is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Error("FindAll is not same as expected")
	}
}

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected domain.Paper
	var err error
	id := 0

	mockSample := mock.NewMockPaperRepository(ctrl)
	mockSample.EXPECT().Store(expected).Return(id, err)

	paperInteractor := NewPaperInteractor(mockSample)
	err = paperInteractor.Add(expected)

	if err != nil {
		t.Error("Store is not same as expected")
	}
}
