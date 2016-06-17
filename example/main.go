package main

import (
	"fmt"
	"net/http"

	"github.com/johanliu/Vidar"
	"github.com/johanliu/Vidar/context"
	"github.com/johanliu/Vidar/middlewares"
	"github.com/johanliu/Vidar/utils"
)

func indexHandler(c *Context) {
	result := map[string]string{"version": "hello world!"}

	c := &context.Context{Response: context.Response{ResponseWriter: w}, Request: r}
	c.JSON(202, result)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page Not Found!")
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
