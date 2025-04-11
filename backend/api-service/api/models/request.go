package models

type RequestData struct {
	Headline string `json:"headline"`
	Url      string `json:"url"`
	Content  string `json:"content"`
}
