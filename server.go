package main

import (
	"net/http"
	"time"
	"github.com/labstack/echo"
)

type PaginatedDocuments struct {
	Page int `json:"page"`
	PerPage int `json:"per_page"`
	Total int `json:"total"`
	Results []Document `json:"results"`
}

type Document struct {
	Id string `json:"id"`
	VersionId string `json:"version_id"`
	Title string `json:"title"`
	DateAdded time.Time  `json:"date_added"`
}


func getDocuments(c echo.Context) error {
	return c.JSON(http.StatusOK, PaginatedDocuments{
		Page: 1,
		PerPage: 10,
		Total: 2,
		Results: []Document{
			Document{Id: "guid1", VersionId: "guid2", Title: "Doc1", DateAdded: time.Now()},
			Document{Id: "guid3", VersionId: "guid4", Title: "Doc2", DateAdded: time.Now()},
		},
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/v1/documents", getDocuments)
	e.Logger.Fatal(e.Start(":1323"))
}

