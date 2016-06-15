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

func nameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test only")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page Not Found! BONG SHAKALAKA!")
}

func main() {

	//TODO: read from config file
	port := "8080"
	host := "localhost"

	Wrapper := utils.New(middlewares.LoggingHandler)
	Wrapper = Wrapper.Append(middlewares.RecoverHandler)

	r := router.New()
	r.Add("GET", "/", Wrapper.Wrap(defaultHandler))
	r.Add("POST", "/", Wrapper.Wrap(nameHandler))
	r.NotFound = Wrapper.Wrap(NotFoundHandler)

	logger.Info.Printf("Running on %s:%s", host, port)
	err := http.ListenAndServe(host+":"+port, r)
	if err != nil {
		logger.Error.Printf("Server start failed: %v", err)
	}
}
