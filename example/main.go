package main

import (
	"github.com/johanliu/Vidar"
	"github.com/johanliu/Vidar/context"
	"github.com/johanliu/Vidar/middlewares"
	"github.com/johanliu/Vidar/utils"
)

func indexHandler(c *context.Context) {
	name := c.Form("name")
	// name := c.Form("name", "jay", "johan")
	result := map[string]string{"message": "hello" + name[0]}
	c.JSON(202, result)
}

func NotFoundHandler(c *context.Context) {
	c.Text(404, "Page Not Found")
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
