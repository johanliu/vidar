# Vidar
A lightweight Golang web framework


## Getting Started

~~~ go
package main

import (
	"fmt"
	"net/http"

	"github.com/johanliu/Vidar/logger"
	"github.com/johanliu/Vidar/middlewares"
	"github.com/johanliu/Vidar/router"
	"github.com/johanliu/Vidar/utils"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page Not Found! BONG SHAKALAKA!")
}

func main() {

	port := "8080"
	host := "localhost"

	Wrapper := utils.New(middlewares.LoggingHandler)
	Wrapper = Wrapper.Append(middlewares.RecoverHandler)

	r := router.New()
	r.Add("GET", "/", Wrapper.Wrap(defaultHandler))
	r.Add("POST", "/", Wrapper.Wrap(defaultHandler))
	r.NotFound = Wrapper.Wrap(NotFoundHandler)

	logger.Info.Printf("Running on %s:%s", host, port)
	err := http.ListenAndServe(host+":"+port, r)
	if err != nil {
		logger.Error.Printf("Server start failed: %v", err)
	}
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
6. Internal TCP/IP Server(optional): Built-in HTTP development server and support for others.
7. Templates(optional): Fast and golang built-in template engine and support for popular templates.
