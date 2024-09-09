package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/a-h/templ"
	helloTemplates "github.com/edwincarlflores/go-htmx/templates/hello"
	pollingTemplates "github.com/edwincarlflores/go-htmx/templates/polling"
	"github.com/edwincarlflores/go-htmx/types"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Render(c echo.Context, statusCode int, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	c.Response().WriteHeader(statusCode)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

func AppTempl() {
	e := echo.New()
	e.Use(middleware.Logger())

	// If you want to serve static assets, just add "static" (or anything you like to name it)
	// directory on the root directory
	// e.Static("/static", "static")

	e.GET("/demo/hello", func(c echo.Context) error {
		return Render(c, http.StatusOK, helloTemplates.HelloPage())
	})
	e.POST("/demo/hello", func(c echo.Context) error {
		name := c.FormValue("name")
		return Render(c, http.StatusOK, helloTemplates.HelloName(name))
	})

	e.GET("/demo/polling", func(c echo.Context) error {
		return Render(c, http.StatusOK, pollingTemplates.PollingPage())
	})

	e.GET("/demo/quotes", func(c echo.Context) error {
		resp, err := http.Get("https://dummyjson.com/quotes/random/10")
		if err != nil {
			log.Fatalf("Error fetching data: %v", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		var quotes []types.Quote
		err = json.Unmarshal(body, &quotes)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
		return Render(c, http.StatusOK, pollingTemplates.Quotes(quotes))
	})

	e.Logger.Fatal(e.Start(":8085"))
}
