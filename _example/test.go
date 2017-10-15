package main

import (
	"github.com/johanliu/mlog"
	"github.com/johanliu/vidar"
	"github.com/johanliu/vidar/plugins"
)

var log = mlog.NewLogger()

type response struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func indexHandler(c *vidar.Context) {
	if c.Method() == "GET" {
		// version := c.QueryParam("version", "123")

		result := `
		<html>
		<header>This is header</header>
		<body>This is body</body>
		</html>
		`

		c.HTML(202, result)
	}

	if c.Method() == "POST" {
		/*
			name := c.FormParam("name")
			value := c.FormParam("value")

			result := map[string]string{
				"message": "hello " + name + value,
			}
			c.JSON(200, result)
		*/
		body := c.Body()
		c.Text(200, string(body[:]))
	}
}

func jsonHandler(c *vidar.Context) {
	res := new(response)

	jp := vidar.Parsers["JSON"]
	if err := jp.Parse(res, c.Request()); err != nil {
		c.Error(vidar.BadRequestError)
	} else {
		c.JSON(200, res)
	}
}

func userHandler(c *vidar.Context) {
	id := c.PathParam("id")
	c.Text(200, id)
}

func notFoundHandler(c *vidar.Context) {
	c.Error(vidar.NotFoundError)
}

func main() {
	v := vidar.New()

	p := vidar.NewPlugin()
	p.Append(plugins.LoggingHandler)
	p.Append(plugins.RecoverHandler)

	v.Router.Add("GET", "/", p.Apply(indexHandler))
	v.Router.Add("POST", "/", p.Apply(indexHandler))
	v.Router.POST("/json", p.Apply(jsonHandler))

	v.Router.NotFound = p.Apply(notFoundHandler)

	v.Run()
}
