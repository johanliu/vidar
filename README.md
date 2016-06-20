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
	result := map[string]string{
		"version" : "0.0.1",
		"name" : "Vidar"
	}

	c.JSON(200, result)
}

func NotFoundHandler(c *Context) {
	c.Text(404, "Page Not Found!")
}

func main() {

	commonHandler := utils.New(middlewares.LoggingHandler)
	commonHandler.Append(middlewares.RecoverHandler)

	v := vidar.New()

	v.Route.GET("/", commonHandler.Wrap(indexHandler))
	v.Route.POST("/", commonHandler.Wrap(indexHandler))
	v.Route.NotFound = commonHandler.Wrap(NotFoundHandler)

	v.Run()
}

~~~


## TODO Lists

1. Routing: Requests to function-call mapping with support for clean and dynamic URLs.
2. Logger: Log the STDOUT and STDERROR messages based on DEBUG, INFO, WARNING, ERROR level.
3. Context: Share request information and others between handler by each request session.
4. Middlewares: Various processing middlewares used by application-layer which support customization.
    - Timing/Logging
    - Recover
    - Parameter checker
    - Request/Response transformation
5. Utilities: Convenient tools to form data, file uploads, cookies, sessions, headers and other HTTP-related metadata.
6. TCP/IP Server(optional): Built-in development server and support for cgi and fastcgi server.
7. Templates(optional): Fast and golang built-in template engine and support for popular templates.
