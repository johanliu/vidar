# Vidar
A lightweight Golang web framework


## Getting Started

~~~ go
package main

import (
    "github.com/johanliu/Vidar"
    "github.com/johanliu/Vidar/context"
    "github.com/johanliu/Vidar/middlewares"
    "github.com/johanliu/Vidar/utils"
)

func indexHandler(c *Context) {
    name := c.Form("name")

    // default value getter
    // name := c.Form("name", "jay", "johan")

	result := map[string]string{
		"version" : "0.0.1",
		"name" : name
	}

	c.JSON(200, result)
}

func personHandler(c *Context) {
    name := c.Params("name")
    c.Text("Hello %s", name)
}

func NotFoundHandler(c *Context) {
	c.Text(404, "Page Not Found!")
}

func main() {
	commonHandler := utils.New(middlewares.LoggingHandler)
	commonHandler.Append(middlewares.AuthHandler)
	commonHandler.Append(middlewares.RecoverHandler)

	v := vidar.New()

	v.Route.GET("/", commonHandler.Wrap(indexHandler))
	v.Route.POST("/person/:name", commonHandler.Wrap(personHandler))
	v.Route.NotFound = commonHandler.Wrap(NotFoundHandler)

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
