package server

import (
	"fmt"
	"github.com/josephshih13/short-url/base62"
	"github.com/josephshih13/short-url/redis"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/http"
	"strings"
)

func InitServer() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.New("web-template").Parse(web_template)),
	}
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		data := template_input{false, "coool", false}

		return c.Render(http.StatusOK, "index", data)
	})

	e.GET("/:id", redirect_url)

	e.POST("/", post_action)

	e.Logger.Fatal(e.Start(":1323"))
}

func base_http(c echo.Context) error {
	data := template_input{false, "coool", false}
	return c.Render(http.StatusOK, "index", data)
}

func redirect_url(c echo.Context) error {
	id := c.Param("id")
	url, err := redis.Get(id)
	if err != nil {
		return base_http(c)
	}
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func invalid_url(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return true
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return true
	}

	return false
}

func post_action(c echo.Context) error {

	url := c.FormValue("url")
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "https://" + url
	}
	print(url)
	if invalid_url(url) {
		fmt.Println("invalid_url")
		return base_http(c)
	}

	idx64, err := redis.Incr("url-index")
	if err != nil {
		fmt.Println("incr fail")
		fmt.Println(err)
		return base_http(c)
	}
	web_id := int(idx64)

	base62_idx := base62.Encode(web_id)
	_ = redis.Set(base62_idx, url)
	full_url = "http://short.josephtest.net/" + base62_idx

	data := template_input{true, full_url, false}

	return c.Render(http.StatusOK, "index", data)
}

type template_input struct {
	Generated bool
	Url       string
	BadUrl    bool
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

const web_template = `{{define "index"}}<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>My short url</title>
</head>
<body>
  <form action="/" method="post">
  <label for="url">Url need to bo shorten :</label>
  <input type="text" id="url" name="url"><br>
  <input type="submit" value="Submit">
</form>
{{if .Generated}}
<p> short url generated : {{.Url}} </p>
{{end}}
{{if .BadUrl}}
<p> bad url </p>
{{end}}

</body>
</html>{{end}}
`
