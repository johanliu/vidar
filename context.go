package vidar

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/johanliu/Vidar/constant"
)

const defaultMaxMemory = 32 << 20 //32MB

// type Parameters []parameter

type Parameters struct {
	key   string
	value map[int]string
}

type Context struct {
	Request    *http.Request
	response   *Response
	parameters *Parameters
	// values     url.Values
	container map[string]interface{}
	status    int
	path      string
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request:    r,
		response:   &Response{ResponseWriter: w},
		parameters: &Parameters{key: "pathParam", value: r.Context().Value("abc").(map[int]string)},
	}
}

// Internal

func (ctx *Context) Set(key string, value interface{}) error {
	if ctx.container == nil {
		ctx.container = make(map[string]interface{})
	}
	ctx.container[key] = value

	return nil
}

func (ctx *Context) Get(key string) (interface{}, bool) {
	if ctx.container != nil {
		value, exist := ctx.container[key]
		return value, exist
	}

	return nil, false
}

// Request

// Query parameters which pick up k-v pairs from URI queries
// for example: http://localhost:8080/index/department?users=alice
// return {"users":"alice"}
func (ctx *Context) QueryParam(key string, defaultvalues ...string) string {
	values, exists := ctx.Request.URL.Query()[key]

	if exists && len(values) > 0 {
		return values[0]
	} else {
		//TODO: index out of range
		return defaultvalues[0]
	}
}

func (ctx *Context) QueryParams() url.Values {
	return ctx.Request.URL.Query()
}

// Path parameters which pick up k-v pairs from URI pathes
// for example: http://localhost:8080/index/department/users/alice
// return {"users":"alice"}
// It should be supported by users which defined router for query like
// "/index/department/users/:users"
func (ctx *Context) PathParam(key string, defaultValues ...string) string {
	value, ok := ctx.getPathParam(key)
	if ok {
		if len(defaultValues) > 0 {
			return defaultValues[0]
		}
	}
	return value
}

func (ctx *Context) getPathParam(key string) (string, bool) {
	//TODO
	return ctx.parameters.value[2], true
}

func (ctx *Context) FormParam(key string, defaultValues ...string) string {
	value := ctx.Request.FormValue(key)
	if value != "" {
		if len(defaultValues) > 0 {
			return defaultValues[0]
		}
	}
	return value
}

func (ctx *Context) FormParams() (url.Values, error) {
	if err := ctx.Request.ParseMultipartForm(defaultMaxMemory); err != nil {
		return nil, err
	}
	return ctx.Request.Form, nil
}

func (ctx *Context) ContentType() string {
	return ctx.Request.Header.Get("Content-Type")
}

func (ctx *Context) Host() string {
	return ctx.Request.Header.Get("Host")
}

// fn.Open()
func (ctx *Context) FormFile(filename string) (*multipart.FileHeader, error) {
	_, fh, err := ctx.Request.FormFile(filename)
	return fh, err
}

func (ctx *Context) Path() string {
	return ctx.path
}

func (ctx *Context) SetPath(path string) {
	ctx.path = path
}

func (ctx *Context) Method() string {
	return ctx.Request.Method
}

//Response

func (ctx *Context) Redirect(code int, url string) {
	if code < 300 || code > 308 {
		log.Error("InvalidRedirectError")
	}

	ctx.response.SetHeader(constant.HeaderLocation, url)
	ctx.response.SetStatus(code)
}

func (ctx *Context) XML(code int, body interface{}) {
	ctx.response.SetHeader(constant.HeaderContentType, constant.MIMEApplicationXMLCharsetUTF8)
	ctx.response.SetStatus(code)

	if err := json.NewEncoder(ctx.response.ResponseWriter).Encode(body); err != nil {
		log.Error("Set payload failed: %v", err)
	}
}

func (ctx *Context) JSON(code int, body interface{}) {
	ctx.response.SetHeader(constant.HeaderContentType, constant.MIMEApplicationJSONCharsetUTF8)
	ctx.response.SetStatus(code)

	if err := json.NewEncoder(ctx.response.ResponseWriter).Encode(body); err != nil {
		log.Error("Set payload failed: %v", err)
	}
}

func (ctx *Context) Text(code int, str string, params ...interface{}) {
	ctx.response.SetHeader(constant.HeaderContentType, constant.MIMETextPlainCharsetUTF8)
	ctx.response.SetStatus(code)

	if _, err := fmt.Fprintf(ctx.response.ResponseWriter, str, params...); err != nil {
		log.Error("Set payload failed: %v", err)
	}
}

func (ctx *Context) HTML(code int, str string, params ...interface{}) {
	ctx.response.SetHeader(constant.HeaderContentType, constant.MIMETextHTMLCharsetUTF8)
	ctx.response.SetStatus(code)

	if _, err := fmt.Fprintf(ctx.response.ResponseWriter, str, params...); err != nil {
		log.Error("Set payload failed: %v", err)
	}

}
