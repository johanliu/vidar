package main

import (
	"github.com/johanliu/Vidar"
	"github.com/johanliu/Vidar/middlewares"
)

func indexHandler(c *vidar.Context) {

	if c.Method() == "GET" {
		version := c.Query("version")

		result := map[string]string{
			"message": "hello " + version,
		}
		c.JSON(202, result)
	}

	if c.Method() == "POST" {
		name := c.FormValue("name")
		value := c.FormValue("value")
		c.Text(200, name+value)
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
	// commonHandler.Append(middlewares.RecoverHandler)

	v := vidar.New()

	// Logical handler for each path
	v.Router.Add("GET", "/main", commonHandler.Use(indexHandler))
	v.Router.Add("POST", "/main", commonHandler.Use(indexHandler))
	v.Router.Add("GET", "/users/:id", commonHandler.Use(indexHandler))
	v.Router.NotFound = commonHandler.Use(notFoundHandler)

	v.Run()
}
