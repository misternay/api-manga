package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/misternay/manga/mangas"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/chapter-all", allMangaHandler)
	e.GET("/chapter-list", allChapterHandler)
	e.GET("/chapter-image", chapterImageHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func allMangaHandler(c echo.Context) error {
	resp := mangas.CallGetAllManga(c)
	return c.JSON(http.StatusOK, resp)
}

func allChapterHandler(c echo.Context) error {
	title := c.QueryParam("title")
	resp := mangas.CallGetAllChapter(c, title)
	return c.JSON(http.StatusOK, resp)
}

func chapterImageHandler(c echo.Context) error {
	title := c.QueryParam("title")
	id := c.QueryParam("id")
	resp := mangas.CallGetAllImageChapter(c, title, id)
	return c.JSON(http.StatusOK, resp)
}
