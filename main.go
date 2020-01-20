package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dannyvankooten/extemplate"
	"github.com/labstack/echo/v4"
)

// Template ...
type Template struct {
	xt *extemplate.Extemplate
}

// NewTemplate ...
func NewTemplate(rootPath string) (*Template, error) {

	tpl := new(Template)
	tpl.xt = extemplate.New()

	err := tpl.xt.ParseDir(rootPath, []string{".html"})
	if err != nil {
		return nil, err
	}

	return tpl, nil
}

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.xt.ExecuteTemplate(w, name, data)
	if err != nil {
		fmt.Println("error", err)
		return err
	}

	return nil
}

// Page ...
type Page struct {
	Name    string
	Content string
}

// Hello ...
func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello.html", "Guys")
}

// About ...
func About(c echo.Context) error {

	data := Page{
		Name: "About", Content: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
Nisi scelerisque eu ultrices vitae auctor eu augue ut lectus. Quam pellentesque nec nam aliquam sem et tortor consequat. 
Pharetra vel turpis nunc eget lorem dolor. Vitae turpis massa sed elementum tempus egestas. Turpis egestas pretium aenean pharetra magna ac placerat. 
Neque ornare aenean euismod elementum nisi quis eleifend quam. In fermentum et sollicitudin ac orci. Ut porttitor leo a diam sollicitudin tempor id eu nisl. 
Sed viverra tellus in hac habitasse platea dictumst vestibulum rhoncus. Lorem ipsum dolor sit amet. A diam sollicitudin tempor id eu. Sit amet facilisis magna etiam. 
Praesent tristique magna sit amet purus gravida.`,
	}
	return c.Render(http.StatusOK, "about.html", data)
}

func main() {

	xt, err := NewTemplate("templates/")
	if err != nil {
		fmt.Println("ErrorExt: ", err)
		os.Exit(1)
	}

	e := echo.New()
	e.Renderer = xt
	e.GET("/hello", Hello)
	e.GET("/about", About)

	e.Logger.Fatal(e.Start(":4000"))

}
