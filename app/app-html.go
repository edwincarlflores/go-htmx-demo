package app

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("templates/html/*.html")),
	}
}

type Count struct {
	Count int
}

func AppHtml() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	count0 := Count{Count: 0}
	e.GET("/demo/count", func(c echo.Context) error {
		count0.Count++
		return c.Render(200, "index", count0)
	})

	counterbody := Count{Count: 0}

	e.GET("/demo/counterbody", func(c echo.Context) error {
		return c.Render(200, "counterbody", counterbody)
	})

	e.POST("/demo/counterbody/count", func(c echo.Context) error {
		counterbody.Count++
		return c.Render(200, "counterbody", counterbody)
	})

	count1 := Count{Count: 0}

	e.GET("/demo/counter", func(c echo.Context) error {
		return c.Render(200, "counter", count1)
	})

	e.POST("/demo/counter/count", func(c echo.Context) error {
		count1.Count++
		return c.Render(200, "count", count1)
	})

	e.Logger.Fatal(e.Start(":8085"))
}
