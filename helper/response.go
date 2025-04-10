package helper

import "todo_api/models"

type ApiResponse struct {
	Status int
	Message string
	Data interface{}
}

type UserResponse struct {
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ToRegisterResponse(user models.User) UserResponse {
	return UserResponse{
		Name: user.Name,
		Username: user.Username,
		Email: user.Email,
	}
}