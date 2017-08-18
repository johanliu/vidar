# Vidar
A lightweight Golang web framework


## Getting Started

~~~ go
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
		c.XML(200, res)
		//c.JSON(200, res)
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
~~~


## TODO Lists

1. Routing: Requests to function-call mapping with support for clean and dynamic URLs.
2. Logger: Log the STDOUT and STDERROR messages based on DEBUG, INFO, WARNING, ERROR level.
3. Context: Process Request/Response information by each request session.
4. Middlewares: Various processing middlewares used by application-layer which support customization.
    - Timing/Logging
    - Recover
    - Parameter checker
    - etc
5. Utilities: Convenient tools to form data, file uploads, cookies, sessions, headers and other HTTP-related metadata.
6. TCP/IP Server(optional): Built-in development server and support for cgi and fastcgi server.
7. Templates(optional): Fast and golang built-in template engine and support for popular templates.
