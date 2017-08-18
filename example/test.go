package main

import (
	"fmt"

	"github.com/johanliu/Vidar"
	"github.com/johanliu/Vidar/middlewares"
	"github.com/johanliu/mlog"
)

var log = mlog.NewLogger()

type response struct {
	name string `json:"name"`
	age  int    `json:"int"`
}

func indexHandler(c *vidar.Context) {
	if c.Method() == "GET" {
		version := c.QueryParam("version")

		result := map[string]string{
			"message": "hello " + version,
		}
		c.JSON(202, result)
	}

	if c.Method() == "POST" {
		name := c.FormParam("name")
		value := c.FormParam("value")
		c.Text(200, name+value)
	}
}

func jsonHandler(c *vidar.Context) {
	res := new(response)

	jp := c.NewParser("JSON")

	if err := jp.Parse(res, c.Request); err != nil {
		log.Info(err.Error())
		c.Text(415, "Parser Error")
	} else {
		c.Text(200, fmt.Sprintf("name: %s, age: %d", res.name, res.age))
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
	v.Router.Add("GET", "/main", commonHandler.Use(indexHandler))
	v.Router.Add("POST", "/main", commonHandler.Use(indexHandler))
	v.Router.POST("/json", commonHandler.Use(jsonHandler))
	v.Router.Add("GET", "/users/:id", commonHandler.Use(indexHandler))
	v.Router.NotFound = commonHandler.Use(notFoundHandler)

	v.Run()
}
