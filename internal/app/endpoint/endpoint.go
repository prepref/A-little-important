package endpoint

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	s Service
}

type Service interface {
	Output() (interface{}, interface{}, error)
	Save(string, string, string, string) error
	ValidateLanguage(string, string, string) (string, error)
	ValidateAdmin(string, string) (bool, error)
	Delete(string, string) error
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Page(ctx echo.Context) error {
	quoteArrayRu, _, err := e.s.Output()
	if err != nil {
		return err
	}
	errRender := ctx.Render(http.StatusOK, "index", quoteArrayRu)
	if errRender != nil {
		return errRender
	}

	return nil
}

func (e *Endpoint) PageEn(ctx echo.Context) error {
	_, quoteArrayEn, err := e.s.Output()
	if err != nil {
		return err
	}
	errRender := ctx.Render(http.StatusOK, "indexEn", quoteArrayEn)
	if errRender != nil {
		return errRender
	}

	return nil
}

func (e *Endpoint) Admin(ctx echo.Context) error {
	quoteArrayRu, _, err := e.s.Output()
	if err != nil {
		return err
	}
	err = ctx.Render(http.StatusOK, "admin", quoteArrayRu)
	if err != nil {
		return err
	}
	return nil
}

func (e *Endpoint) AdminEn(ctx echo.Context) error {
	_, quoteArrayEn, err := e.s.Output()
	if err != nil {
		return err
	}
	err = ctx.Render(http.StatusOK, "adminEn", quoteArrayEn)
	if err != nil {
		return err
	}
	return nil
}

func (e *Endpoint) AdminDelete(ctx echo.Context) error {
	login := ctx.FormValue("login")
	password := ctx.FormValue("password")
	author := ctx.FormValue("author")
	quote := ctx.FormValue("quote")
	f, err := e.s.ValidateAdmin(login, password)
	if err != nil {
		return err
	}
	if f {
		err := e.s.Delete(author, quote)
		if err != nil {
			return err
		}
		err = ctx.Render(http.StatusOK, "adminDelete.html", nil)
		if err != nil {
			return err
		}
	} else {
		err := ctx.Render(http.StatusOK, "adminWrong.html", nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Endpoint) AdminDeleteEn(ctx echo.Context) error {
	login := ctx.FormValue("login")
	password := ctx.FormValue("password")
	author := ctx.FormValue("author")
	quote := ctx.FormValue("quote")
	f, err := e.s.ValidateAdmin(login, password)
	if err != nil {
		return err
	}
	if f {
		err := e.s.Delete(author, quote)
		if err != nil {
			return err
		}
		err = ctx.Render(http.StatusOK, "adminDeleteEn.html", nil)
		if err != nil {
			return err
		}
	} else {
		err := ctx.Render(http.StatusOK, "adminWrongEn.html", nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Endpoint) CreateQuote(ctx echo.Context) error {
	err := ctx.Render(http.StatusOK, "create.html", nil)
	if err != nil {
		return err
	}
	return nil
}

func (e *Endpoint) CreateQuoteEn(ctx echo.Context) error {
	err := ctx.Render(http.StatusOK, "createEn.html", nil)
	if err != nil {
		return err
	}
	return nil
}

func (e *Endpoint) Data(ctx echo.Context) error {
	author := ctx.FormValue("author")
	quote := ctx.FormValue("quote")
	name := ctx.FormValue("name")
	if author == "" || quote == "" || name == "" {
		err := ctx.Render(http.StatusOK, "no-data.html", nil)
		if err != nil {
			return err
		}
	} else {
		language, err := e.s.ValidateLanguage(author, quote, name)
		if err != nil {
			return err
		}
		if language == "en" {
			err := ctx.Render(http.StatusOK, "no-data.html", nil)
			if err != nil {
				return err
			}
		} else {
			err = e.s.Save(author, name, quote, language)
			if err != nil {
				return err
			}
			err = ctx.Render(http.StatusOK, "with-data.html", nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *Endpoint) DataEn(ctx echo.Context) error {
	author := ctx.FormValue("author")
	quote := ctx.FormValue("quote")
	name := ctx.FormValue("name")
	if author == "" || quote == "" || name == "" {
		err := ctx.Render(http.StatusOK, "no-dataEn.html", nil)
		if err != nil {
			return err
		}
	} else {
		language, err := e.s.ValidateLanguage(author, quote, name)
		if err != nil {
			return err
		}
		if language == "ru" {
			err := ctx.Render(http.StatusOK, "no-dataEn.html", nil)
			if err != nil {
				return err
			}
		} else {
			err = e.s.Save(author, name, quote, language)
			if err != nil {
				return err
			}

			err = ctx.Render(http.StatusOK, "with-dataEn.html", nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
