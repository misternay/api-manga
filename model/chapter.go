package model

type Chapter struct {
	ID      string `json:"id"`
	Chapter string `json:"chapter"`
	Credit  string `json:"credit"`
	Date    string `json:"date"`
}

type ChapterImage struct {
	ID     string `json:"id"`
	ImgURL string `json:"imgURL"`
}
