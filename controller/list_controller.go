package controller

import (
	"net/http"
	"strconv"
	"todo_api/helper"
	"todo_api/services"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ListController interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error 
	FindByUserId(c echo.Context) error
	//FindId(c echo.Context) error
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

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	listPayload.UserId = userId

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

	c.Response().Header().Add("Content-Type", "application/json")

	return c.JSON(http.StatusOK, apiResponse)
}

// Update implements ListController.
func (controller *ListControllerImpl) Update(c echo.Context) error {
	listPayload := new(helper.ListRequestUpdate)
	err := c.Bind(listPayload)
	if err != nil {
		panic(err)
	}
	
	listId := c.Param("list_id")
	id, err := strconv.Atoi(listId)
	if err != nil{
		return c.JSON(http.StatusBadRequest, helper.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "ID tidak valid",
		})
	}
	listPayload.Id = id

	result, err := controller.ListService.Update(c, *listPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Gagal Mengubah Data"})
	}

	apiResponse := helper.ApiResponse{
		Status:  http.StatusOK,
		Message: "Berhasil Ubah Data",
		Data:    result,
	}

	c.Response().Header().Add("Content-Type", "application/json")

	return c.JSON(http.StatusOK, apiResponse)
}

// Delete implements ListController.
func (controller *ListControllerImpl) Delete(c echo.Context) error  {
	listId := c.Param("list_id")
	id, err := strconv.Atoi(listId)
	if err != nil{
		return c.JSON(http.StatusBadRequest, helper.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "ID tidak valid",
		})
	}

	controller.ListService.Delete(c, id)
	apiResponse := helper.ApiResponse{
		Status:  http.StatusOK,
		Message: "Berhasil Hapus Data",
	}

	return c.JSON(http.StatusOK, apiResponse)
}

// FindByUserId implements ListController.
func (controller *ListControllerImpl) FindByUserId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	list, err := controller.ListService.FindByUserId(c, int(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, list)
}

// FindId implements ListController.
// func (controller *ListControllerImpl) FindId(c echo.Context) error {
// 	listId := c.Param("list_id")
// 	id, err := strconv.Atoi(listId)
// 	if err != nil{
// 		return c.JSON(http.StatusBadRequest, helper.ApiResponse{
// 			Status:  http.StatusBadRequest,
// 			Message: "ID tidak valid",
// 		})
// 	}

// 	controller.ListService.
// }

func NewListControllerImpl(listService services.ListService) ListController {
	return &ListControllerImpl{
		ListService: listService,
	}
}
