package main

import (
	"flag"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bryanl/websay/pkg/say"
)

func main() {
	var bgColor string
	flag.StringVar(&bgColor, "bg-color", "white", "background color")

	var fgColor string
	flag.StringVar(&fgColor, "fg-color", "black", "foreground color")

	var message string
	flag.StringVar(&message, "message", "Hello", "message")

	var sayType string
	flag.StringVar(&sayType, "type", "default", "type")

	flag.Parse()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		content, err := say.Say(message, bgColor, fgColor, sayType)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.HTML(http.StatusOK, content)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
