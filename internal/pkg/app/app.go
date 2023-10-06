package app

import (
	"html/template"
	"io"
	"log"
	"quotes-project/internal/app/endpoint"
	"quotes-project/internal/app/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func New() (*App, error) {
	a := &App{}
	t := &Template{
		templates: template.Must(template.ParseGlob("web/templates/*.html")),
	}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()
	a.echo.Renderer = t
	a.echo.Static("/static", "web/static")
	a.echo.GET("/", a.e.Page)
	a.echo.GET("/en", a.e.PageEn)
	a.echo.GET("/create", a.e.CreateQuote)
	a.echo.GET("/createEn", a.e.CreateQuoteEn)
	a.echo.GET("/save-quotes", a.e.Data)
	a.echo.POST("/save-quotes", a.e.Data)
	a.echo.GET("/save-quotesEn", a.e.DataEn)
	a.echo.POST("/save-quotesEn", a.e.DataEn)
	a.echo.GET("/admin", a.e.Admin)
	a.echo.GET("/delete", a.e.AdminDelete)
	a.echo.POST("/delete", a.e.AdminDelete)
	a.echo.GET("/adminEn", a.e.AdminEn)
	a.echo.GET("/deleteEn", a.e.AdminDeleteEn)
	a.echo.POST("/deleteEn", a.e.AdminDeleteEn)
	return a, nil
}

func (a *App) Run() error {
	log.Println("server running")
	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
