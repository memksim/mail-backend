package model

type Mail struct {
	Id         int64  `json:"id"`
	Sender     User   `json:"sender"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	IsBookmark bool   `json:"is_bookmark"`
	Time       int64  `json:"time"`
	IsRead     bool   `json:"is_read"`
}
