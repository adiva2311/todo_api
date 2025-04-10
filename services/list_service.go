package services

import (
	"database/sql"
	"errors"
	"fmt"
	"todo_api/helper"
	"todo_api/models"
	"todo_api/repositories"

	"github.com/labstack/echo/v4"
)

type ListService interface {
	Create(c echo.Context, request helper.ListRequestCreate) (helper.ListResponse, error)
	Update(c echo.Context, request helper.ListRequestUpdate) (helper.ListResponse, error)
	Delete(c echo.Context, listId uint, userId int) error
	FindByUserId(c echo.Context, userId int) ([]models.List, error)
}

type ListServiceImpl struct {
	ListRepo repositories.ListRepository
	Db       *sql.DB
}

// Create implements ListService.
func (service *ListServiceImpl) Create(c echo.Context, request helper.ListRequestCreate) (helper.ListResponse, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	userIdFloat, ok := c.Get("user_id").(float64)
	if !ok {
		return helper.ListResponse{}, errors.New("invalid user_id type in context")
	}
	userId := int(userIdFloat)
	list := models.List{
		Title:       request.Title,
		Information: request.Information,
		Completed:   request.Complete,
		UserId:      userId,
	}

	fmt.Println("Title masuk:", request.Title)
	fmt.Println("Information masuk:", request.Information)

	savedList, err := service.ListRepo.Create(c, tx, list)
	if err != nil {
		panic(err)
	}
	return helper.ToListResponse(savedList), nil
}

// Update implements ListService.
func (service *ListServiceImpl) Update(c echo.Context, request helper.ListRequestUpdate) (helper.ListResponse, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	userIdFloat, ok := c.Get("user_id").(float64)
	if !ok {
		return helper.ListResponse{}, errors.New("invalid user_id type in context")
	}
	userId := int(userIdFloat)

	_, err = service.ListRepo.FindId(c, tx, uint(request.Id), userId)
	if err != nil {
		return helper.ListResponse{}, errors.New("list_id not found")
	}

	list := models.List{
		Id:          uint(request.Id),
		Title:       request.Title,
		Information: request.Information,
		Completed:   request.Complete,
		UserId:      userId,
	}

	updatedList, err := service.ListRepo.Update(c, tx, list)
	if err != nil {
		panic(err)
	}

	return helper.ToListResponse(updatedList), nil
}

// Delete implements ListService.
func (service *ListServiceImpl) Delete(c echo.Context, listId uint, userId int) error {
	tx, err := service.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	_, err = service.ListRepo.FindId(c, tx, listId, userId)
	if err != nil {
		return errors.New("list_id not found")
	}

	err = service.ListRepo.Delete(c, tx, listId, userId)
	if err != nil {
		return err
	}

	return nil
}

// FindByUserId implements ListService.
func (service *ListServiceImpl) FindByUserId(c echo.Context, userId int) ([]models.List, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	list, err := service.ListRepo.FindByUserId(c, tx, userId)
	if err != nil {
		panic(err)
	}

	return list, nil
}

func NewListServiceImpl(listRepo repositories.ListRepository, db *sql.DB) ListService {
	return &ListServiceImpl{
		ListRepo: listRepo,
		Db:       db,
	}
}
