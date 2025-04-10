package models

type List struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Information string `json:"information"`
	Complete    bool   `json:"complete"`
	UserId      uint   `json:"user_id"`
}

func (List) TableName() string {
	return "list"
}