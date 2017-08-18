package main

import (
	"github.com/johanliu/Vidar"
	"github.com/johanliu/Vidar/middlewares"
	"github.com/johanliu/mlog"
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

	jp := c.NewParser("JSON")
	if err := jp.Parse(res, c.Request); err != nil {
		c.Text(415, "Parser Error")
	} else {
		//c.XML(200, res)
		c.JSON(200, res)
	}
}

func userHandler(c *vidar.Context) {
	id := c.PathParam("id")
	c.Text(200, id)
}

func notFoundHandler(c *vidar.Context) {
	c.Text(404, "Page Not Found")
}

func main() {
	// Common utils handler for all path defined below
	commonHandler := vidar.NewChain()

	commonHandler.Append(middlewares.LoggingHandler)
	commonHandler.Append(middlewares.RecoverHandler)

	v := vidar.New()

	// Handlers
	v.Router.Add("GET", "/", commonHandler.Use(indexHandler))
	v.Router.Add("POST", "/", commonHandler.Use(indexHandler))
	v.Router.POST("/json", commonHandler.Use(jsonHandler))
	v.Router.Add("GET", "/users/:id", commonHandler.Use(userHandler))
	v.Router.NotFound = commonHandler.Use(notFoundHandler)

	v.Run()
}
