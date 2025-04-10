package helper

import (
	"todo_api/models"
)

type ApiResponse struct {
	Status int
	Message string
	Data interface{}
}

type UserResponse struct {
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
}

func ToRegisterResponse(user models.User) UserResponse {
	return UserResponse{
		Name: user.Name,
		Username: user.Username,
		Email: user.Email,
	}
}

type LoginResponse struct{
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ToLoginResponse(user models.User, token string) LoginResponse {
	return LoginResponse{
		Name: user.Name,
		Username: user.Username,
		Email: user.Email,
		Token: token,
	}
}

type ListResponse struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Information string `json:"information"`
	Complete    bool   `json:"complete"`
	UserId      uint   `json:"user_id"`
}


func ToListResponse(list models.List) ListResponse {
	return ListResponse{
		Id:          list.Id,
		Title:       list.Title,
		Information: list.Information,
		Complete:    list.Complete,
		UserId:      list.UserId,
	}
}