package main

import "net/http"

func main() {

	//TODO
	port := "8080"
	host := "localhost"

	//TODO router

	mux := http.NewServeMux()
	logger.Info.Println("Running on %s:%s", host, port)
	http.ListenAndServe(host+":"+port, mux)
}
