package vidar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	http.ResponseWriter
	status        int
	contentType   string
	contentLength int
}

func (ctx *Context) Write(data []byte) (int, error) {
	if ctx.Response.status == 0 {
		ctx.Response.status = http.StatusOK
	}

	size, err := ctx.ResponseWriter.Write(data)
	ctx.Response.contentLength += size

	return size, err
}

func (ctx *Context) JSON(code int, body interface{}) {
	if err := ctx.SetContentType("application/json; charset=utf-8"); err != nil {
		fmt.Printf("Set content type failed: %v", err)
	}

	if err := ctx.SetStatus(code); err != nil {
		fmt.Printf("Set status code failed: %v", err)
	}

	if err := json.NewEncoder(ctx.ResponseWriter).Encode(body); err != nil {
		fmt.Printf("Set payload failed: %v", err)
	}
}

func (ctx *Context) Text(code int, str string, params ...interface{}) {
	if err := ctx.SetContentType("text/plain; charset=utf-8"); err != nil {
		fmt.Printf("Set content type failed: %v", err)
	}

	if err := ctx.SetStatus(code); err != nil {
		fmt.Printf("Set status code failed: %v", err)
	}

	if len(params) > 0 {
		if _, err := fmt.Fprintf(ctx.ResponseWriter, str, params...); err != nil {
			fmt.Printf("Set payload failed: %v", err)
		}
	} else {
		if _, err := io.WriteString(ctx.ResponseWriter, str); err != nil {
			fmt.Printf("Set payload failed: %v", err)
		}
	}
}

func (response *Response) SetContentType(value string) error {
	return response.SetHeader("Content-Type", value)
}

func (response *Response) SetStatus(code int) error {
	response.ResponseWriter.WriteHeader(code)

	return nil
}

func (response *Response) SetHeader(key string, value string) error {
	if len(value) == 0 {
		response.ResponseWriter.Header().Del(key)
	} else {
		response.ResponseWriter.Header().Set(key, value)
	}

	return nil
}

func (response *Response) Size() (int, error) {
	return response.contentLength, nil
}
