package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"net/http"
)

func main() {

	e := echo.New()
	// the file server for rice. "public" is the folder where the files come from.
	assetHandler := http.FileServer(rice.MustFindBox("public").HTTPBox())
	// serves the index.html from rice
	e.GET("/", echo.WrapHandler(assetHandler))

	// servers other static files
	e.GET("/static/*", echo.WrapHandler(assetHandler))
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
