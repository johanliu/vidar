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
