package models

type User struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type List struct {
	Id uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Complete string `json:"complete"`
	UserId User `json:"user_id"`
}

func (User) TableName() string {
	return "user"
}

func (List) TableName() string {
	return "list"
}