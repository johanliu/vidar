package main

import (
	"fmt"
	"net/http"

	"github.com/johanliu/Vidar/logger"
	"github.com/johanliu/Vidar/middlewares"
	"github.com/johanliu/Vidar/utils"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println(r)

	fmt.Fprintf(w, "Hello World!")
}

func main() {

	//TODO: read from config file
	port := "8080"
	host := "localhost"

	//TODO router should be defined in sole file

	Wrapper := utils.New(middlewares.LoggingHandler)
	Wrapper = Wrapper.Append(middlewares.RecoverHandler)

	http.Handle("/", Wrapper.Wrap(defaultHandler))
	logger.Info.Printf("Running on %s:%s", host, port)
	http.ListenAndServe(host+":"+port, nil)
}
