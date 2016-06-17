# Vidar
A lightweight Golang web framework


## Getting Started

~~~ go
package main

import (
	"fmt"
	"net/http"

	"github.com/johanliu/Vidar"
	"github.com/johanliu/Vidar/middlewares"
	"github.com/johanliu/Vidar/utils"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page Not Found!")
}

func main() {

	commonHandler := utils.New(middlewares.LoggingHandler)
	commonHandler.Append(middlewares.RecoverHandler)

	v := vidar.New()

	v.Route.Add("GET", "/", commonHandler.Wrap(indexHandler))
	v.Route.Add("POST", "/", commonHandler.Wrap(indexHandler))
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
