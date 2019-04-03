package model

type MangaModel struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Chapter string `json:"chapter"`
	Hot     string `json:"_hot"`
	Cover   string `json:"cover"`
}
