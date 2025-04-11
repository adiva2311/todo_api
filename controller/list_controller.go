package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"todo_api/helper"
	"todo_api/services"

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
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Format request tidak valid"})
	}
	fmt.Printf("Bind result: %+v\n", listPayload)

	// Ambil user_id dari context (bukan "user")
	userIdInterface := c.Get("user_id")
	if userIdInterface == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized"})
	}

	userIdFloat, ok := userIdInterface.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid user_id type"})
	}

	listPayload.UserId = int(userIdFloat)
	fmt.Printf("user_id from token: %v\n", listPayload.UserId)

	result, err := controller.ListService.Create(c, *listPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Gagal Menambah Data"})
	}

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
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Format request tidak valid"})
	}

	// Ambil user_id dari context (bukan "user")
	userIdInterface := c.Get("user_id")
	if userIdInterface == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized"})
	}

	userId, ok := userIdInterface.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid user_id type"})
	}
	listPayload.UserId = int(userId)

	listId := c.Param("list_id")
	id, err := strconv.Atoi(listId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "ID tidak valid",
		})
	}
	listPayload.Id = id

	fmt.Printf("user_id from token: %v\n", listPayload.UserId)
	fmt.Printf("Bind result: %+v\n", listPayload)

	result, err := controller.ListService.Update(c, *listPayload)
	if err != nil {
		fmt.Printf("Update error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Gagal Mengubah Data",
		})
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
func (controller *ListControllerImpl) Delete(c echo.Context) error {
	id := c.Param("list_id")
	listId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "ID tidak valid",
		})
	}
	fmt.Printf("list_id : %v\n", listId)

	// Ambil user_id dari context (bukan "user")
	userIdInterface := c.Get("user_id")
	if userIdInterface == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized"})
	}

	userId, ok := userIdInterface.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid user_id type"})
	}
	fmt.Printf("user_id from token: %v\n", userId)

	err = controller.ListService.Delete(c, uint(listId), int(userId))
	if err != nil {
		fmt.Printf("Delete error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Gagal Menghapus Data",
		})
	}

	apiResponse := helper.ApiResponse{
		Status:  http.StatusOK,
		Message: "Berhasil Hapus Data",
	}

	return c.JSON(http.StatusOK, apiResponse)
}

// FindByUserId implements ListController.
func (controller *ListControllerImpl) FindByUserId(c echo.Context) error {
	// Ambil user_id dari context (bukan "user")
	userIdInterface := c.Get("user_id")
	if userIdInterface == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized"})
	}

	userIdFloat, ok := userIdInterface.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid user_id type"})
	}

	UserId := int(userIdFloat)
	fmt.Printf("user_id from token: %v\n", UserId)

	list, err := controller.ListService.FindByUserId(c, UserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	apiResponse := helper.ApiResponse{
		Status:  http.StatusOK,
		Message: "Berhasil Get Data",
		Data:    list,
	}

	c.Response().Header().Add("Content-Type", "application/json")

	return c.JSON(http.StatusOK, apiResponse)
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
