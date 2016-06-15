package main

import (
	"fmt"
	"net/http"

	"github.com/johanliu/Vidar/logger"
	"github.com/johanliu/Vidar/middlewares"
	"github.com/johanliu/Vidar/utils"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {

	//TODO: read from config file
	port := "8080"
	host := "localhost"

	//TODO router should be defined in sole file

	logWrapper := utils.New(middlewares.LoggingHandler)

	http.Handle("/", logWrapper.Wrap(defaultHandler))
	logger.Info.Printf("Running on %s:%s", host, port)
	http.ListenAndServe(host+":"+port, nil)
}
