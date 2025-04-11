package models

type List struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Information string `json:"information"`
	Completed   bool   `json:"completed"`
	UserId      int    `json:"user_id"`
}

func (List) TableName() string {
	return "list"
}
