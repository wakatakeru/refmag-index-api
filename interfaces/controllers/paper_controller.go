package controllers

import (
	"net/http"
	"strconv"

	"github.com/wakatakeru/refmag-index-api/domain"
	"github.com/wakatakeru/refmag-index-api/interfaces/database"
	"github.com/wakatakeru/refmag-index-api/usecase"
)

type PaperController struct {
	Interactor usecase.PaperInteractor
	JWTHandler JWTHandler
}

func NewPaperController(sqlHandler database.SqlHandler, jwtHandler JWTHandler) *PaperController {
	return &PaperController{
		Interactor: usecase.PaperInteractor{
			PaperRepository: &database.PaperRepository{
				SqlHandler: sqlHandler,
			},
		},
		JWTHandler: jwtHandler,
	}
}

func (controller *PaperController) Create(c Context) {
	// TODO: Fix
	_, err := controller.JWTHandler.Verify(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	paper := domain.Paper{}
	err = c.Bind(&paper)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = controller.Interactor.Add(paper)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, paper)
}

func (controller *PaperController) Show(c Context) {
	// TODO: Fix
	_, err := controller.JWTHandler.Verify(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	paper, err := controller.Interactor.Paper(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, paper)
}

func (controller *PaperController) Index(c Context) {
	// TODO: Fix
	_, err := controller.JWTHandler.Verify(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	papers, err := controller.Interactor.Papers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, papers)
}
