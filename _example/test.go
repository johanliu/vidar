package main

import (
	"github.com/johanliu/mlog"
	"github.com/johanliu/vidar"
	_ "github.com/johanliu/vidar/parsers"
	"github.com/johanliu/vidar/plugins"
)

var log = mlog.NewLogger()

type response struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func indexHandler(c *vidar.Context) {
	if c.Method() == "GET" {
		result := `
		<html>
		<header>This is header</header>
		<body>This is body</body>
		</html>
		`

		c.HTML(202, result)
	}

	if c.Method() == "POST" {
		name := c.FormParam("name")
		value := c.FormParam("value")

		result := map[string]string{
			"message": "hello " + name + value,
		}
		c.JSON(200, result)
	}
}

func jsonHandler(c *vidar.Context) {
	res := new(response)

	p := vidar.Parsers["JSON"]
	if err := p.Parse(res, c.Request()); err != nil {
		c.Error(vidar.BadRequestError)
	} else {
		c.JSON(200, res)
	}
}

func staffHandler(c *vidar.Context) {
	username := c.PathParam("username")

	c.JSON(200, map[string]string{"username": username})
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
	p := v.Plugin

	p.Append(plugins.LoggingHandler)
	p.Append(plugins.RecoverHandler)

	v.Router.Add("GET", "/", p.Apply(indexHandler))
	v.Router.Add("POST", "/", p.Apply(indexHandler))

	// Custom parser
	v.Router.POST("/json/read/here", p.Apply(jsonHandler))
	v.Router.GET("/json/read/here", p.Apply(staffHandler))

	// Path parameter
	v.Router.GET("/staff/:username/id", p.Apply(staffHandler))

	// File handler
	//v.Router.File("/", "./public/hello.txt")

	// index.html by default
	v.Router.File("/", "./public")

	// Static resource, serve all files in directory
	v.Router.Static("/portal", "./public")

	// NotFound handler
	v.Router.NotFound = p.Apply(notFoundHandler)

	v.Run()
}
