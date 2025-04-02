package models

type Source struct {
	// For now only one field, but later on can be extended
	Url string `json:"url"`
}

type ResponseData struct {
	Score   int      `json:"score"`
	Sources []Source `json:"sources"`
	Resutl  string   `json:"result"`
}
