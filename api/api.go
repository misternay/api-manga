package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	host          = "http://www.mangastalker-th.com/"
	selectManga   = "SelectFromTableSqlite3110.php"
	selectChapter = "GetChapterFromTableSqlite.php"
	selectImage   = "GetImageFromTableSqlite.php"
)

func GetAllManga() []byte {
	resp, err := http.Get(host + selectManga + "?a=Name&b=date&c=DESC&fbclid=IwAR3-LdBjXxlv4_TsNX_L2MHPllDiSyjNCfpgCm_BVoS-uQFLb4qzx4jlpxc")

	if err != nil {
		println(err)
	}

	body, errs := ioutil.ReadAll(resp.Body)

	if errs != nil {
		println(errs)

	}

	defer resp.Body.Close()
	return body
}

func GetChapterList(title string) []byte {
	resp, err := http.Get(host + selectChapter + "?a=" + removeSpace(title) + "&b=sqlite:MangaDB/" + removeSpace(title) + ".sqlite")

	if err != nil {
		println(err)
	}

	body, errs := ioutil.ReadAll(resp.Body)

	if errs != nil {
		println(errs)

	}

	defer resp.Body.Close()
	return body
}

func removeSpace(value string) string {
	return strings.Join(strings.Fields(value), "")
}

func GetAllImageChapter(title, id string) []byte {
	resp, err := http.Get(host + selectImage + "?a=" + removeSpace(title) + strconv.Itoa(11) + "&b=sqlite:MangaDB/" + removeSpace(title) + ".sqlite")

	if err != nil {
		println(err)
	}

	body, errs := ioutil.ReadAll(resp.Body)
	fmt.Printf("%v", resp.Request.URL)
	if errs != nil {
		println(errs)

	}

	defer resp.Body.Close()
	return body
}
