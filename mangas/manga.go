package mangas

import (
	"encoding/json"

	"github.com/labstack/echo"
	"github.com/misternay/manga/api"
	"github.com/misternay/manga/model"
)

func CallGetAllManga(c echo.Context) []model.MangaModel {
	ps := []model.MangaModel{}
	body := api.GetAllManga()
	error := json.Unmarshal(body, &ps)

	if error != nil {
		c.Logger().Fatal(error)
	}
	return ps
}

func CallGetAllChapter(c echo.Context, title string) []model.Chapter {
	ps := []model.Chapter{}
	body := api.GetChapterList(title)
	err := json.Unmarshal(body, &ps)
	if err != nil {
		c.Logger().Error(err)
	}
	return ps
}

func CallGetAllImageChapter(c echo.Context, title string, id string) []model.ChapterImage {
	ps := []model.ChapterImage{}
	body := api.GetAllImageChapter(title, id)
	err := json.Unmarshal(body, &ps)
	if err != nil {
		c.Logger().Error(err)
	}
	return ps
}
