package controller

import (
	"net/http"
	"todo_api/helper"
	"todo_api/services"

	"github.com/labstack/echo/v4"
)

type ListController interface {
	Create(c echo.Context) error
	Update(c echo.Context) 
	Delete(c echo.Context)
	FindByUserId(c echo.Context)
	FindId(c echo.Context)
}

type ListControllerImpl struct {
	ListService services.ListService
}

// Create implements ListController.
func (controller *ListControllerImpl) Create(c echo.Context) error {
	listPayload := new(helper.ListRequestCreate)
	err := c.Bind(listPayload)
	if err != nil {
		panic(err)
	}

	result, err := controller.ListService.Create(c, *listPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Gagal Menambah Data"})
	}
	// result, err := controller.ListService.Create(c, helper.ListRequestCreate{
	// 	Title: listPayload.Title,
	// 	Information: listPayload.Information,
	// 	Complete: listPayload.Complete,
	// })
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Gagal register", "error": err.Error()})
	// }

	apiResponse := helper.ApiResponse{
		Status:  http.StatusOK,
		Message: "Berhasil Tambah Data",
		Data:    result,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

// Update implements ListController.
func (controller *ListControllerImpl) Update(c echo.Context) {
	panic("unimplemented")
}

// Delete implements ListController.
func (controller *ListControllerImpl) Delete(c echo.Context) {
	panic("unimplemented")
}

// FindByUserId implements ListController.
func (controller *ListControllerImpl) FindByUserId(c echo.Context) {
	panic("unimplemented")
}

// FindId implements ListController.
func (controller *ListControllerImpl) FindId(c echo.Context) {
	panic("unimplemented")
}

func NewListControllerImpl(listService services.ListService) ListController {
	return &ListControllerImpl{
		ListService: listService,
	}
}
