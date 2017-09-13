package main

import (
	"bytes"
	"context"
	"io"

	"github.com/grafana/grafana/pkg/cmd/grafana-cli/logger"
	"github.com/johanliu/Vidar"
	"github.com/johanliu/Vidar/middlewares"
)

/*
POST  HTTP/1.1
Host: localhost:8080?version=12345
Cache-Control: no-cache
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW

----WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="name"

johan
----WebKitFormBoundary7MA4YWxkTrZu0gW

============================================================

POST  HTTP/1.1
Host: localhost:8080?version=12345
Cache-Control: no-cache
Content-Type: application/x-www-form-urlencoded

name=johan
*/

func indexHandler(c *context.Context) {
	version := c.Query("version")
	name := c.Form("name")
	// name := c.Form("name", "jay", "johan")

	result := map[string]string{
		"message": "hello " + name,
		"version": version,
	}
	c.JSON(202, result)
}

/*

POST /file?version=12345 HTTP/1.1
Host: localhost:8080
Cache-Control: no-cache
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW

----WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="age"

18
----WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="pw"; filename=""
Content-Type:


----WebKitFormBoundary7MA4YWxkTrZu0gW

*/

func fileHandler(c *context.Context) {

	var b bytes.Buffer

	file, err := c.File("pw")
	if err != nil {
		logger.Error.Panicf("File read failed: %v", err)
	}
	defer file.Close()

	io.Copy(&b, file)
	c.Text(200, (&b).String())
}

/*
func personHandler(c *context.Context) {
	name := c.Params("name")

	c.Text(200, "Hello %s", name)
}*/

func NotFoundHandler(c *context.Context) {
	c.Text(404, "Page Not Found")
}

func main() {

	commonHandler := utils.New(middlewares.LoggingHandler)
	commonHandler.Append(middlewares.RecoverHandler)

	v := vidar.New()

	v.Route.GET("/", commonHandler.Wrap(indexHandler))
	v.Route.POST("/", commonHandler.Wrap(indexHandler))
	v.Route.POST("/file", commonHandler.Wrap(fileHandler))
	v.Route.NotFound = commonHandler.Wrap(NotFoundHandler)

	v.Run()
}
