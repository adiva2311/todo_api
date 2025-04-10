package helper

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ListRequestCreate struct {
	Title       string `json:"title"`
	Information string `json:"information"`
	Complete    bool   `json:"complete"`
}

type ListRequestUpdate struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Information string `json:"information"`
	Complete    bool   `json:"complete"`
}