package vidar

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/johanliu/mlog"
)

const defaultMaxMemory = 32 << 20 //32MB
const indexPage = "index.html"

// type Parameters []parameter

type Parameters struct {
	key   string
	value map[int]string
}

type Context struct {
	request    *http.Request
	response   *Response
	parameters *Parameters
	// values     url.Values
	container map[string]interface{}
	status    int
	path      string
	Log       *mlog.Logger
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	c := &Context{
		request:  r,
		response: &Response{ResponseWriter: w},
	}

	return c
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

func (ctx *Context) Response() http.ResponseWriter {
	return ctx.response.ResponseWriter
}

func (ctx *Context) Request() *http.Request {
	return ctx.request
}

func (ctx *Context) Call(h http.Handler) {
	if f, ok := h.(ContextFunc); ok {
		f(ctx)
	}
}

// Request

func (ctx *Context) Method() string {
	return ctx.request.Method
}

func (ctx *Context) ContentType() string {
	return ctx.request.Header.Get(HeaderContentType)
}

func (ctx *Context) Body() []byte {
	body, err := ioutil.ReadAll(ctx.request.Body)
	if err != nil {
		ctx.Log.Error(err)
	}

	return body
}

func (ctx *Context) RealIP() string {
	if ip := ctx.request.Header.Get(HeaderXForwardedFor); ip != "" {
		return strings.Split(ip, ",")[0]
	} else if ip := ctx.request.Header.Get(HeaderXRealIP); ip != "" {
		return ip
	}

	host, _, err := net.SplitHostPort(ctx.request.RemoteAddr)
	if err != nil {
		log.Error(err)
	}
	return host
}

func (ctx *Context) Scheme() string {
	if ctx.request.TLS != nil {
		return "HTTPS"
	}
	return "HTTP"
}

// Query parameters which pick up k-v pairs from URI queries
// for example: http://localhost:8080/index/department?users=alice
// return {"users":"alice"}
func (ctx *Context) QueryParam(key string, defaultvalues ...string) string {
	values, exists := ctx.request.URL.Query()[key]

	if exists && len(values) > 0 {
		return values[0]
	} else {
		//TODO: index out of range
		return defaultvalues[0]
	}
}

func (ctx *Context) QueryParams() url.Values {
	return ctx.request.URL.Query()
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
	value := ctx.request.FormValue(key)
	if value != "" {
		if len(defaultValues) > 0 {
			return defaultValues[0]
		}
	}
	return value
}

func (ctx *Context) FormParams() (url.Values, error) {
	if err := ctx.request.ParseForm(); err != nil {
		return nil, err
	}
	return ctx.request.Form, nil
}

func (ctx *Context) MultiFormParams() (url.Values, error) {
	if err := ctx.request.ParseMultipartForm(defaultMaxMemory); err != nil {
		return nil, err
	}
	return ctx.request.Form, nil
}

func (ctx *Context) Cookie(name string) (*http.Cookie, error) {
	return ctx.request.Cookie(name)
}

func (ctx *Context) Cookies() []*http.Cookie {
	return ctx.request.Cookies()
}

// fn.Open()
func (ctx *Context) FormFile(filename string) (*multipart.FileHeader, error) {
	_, fh, err := ctx.request.FormFile(filename)
	return fh, err
}

func (ctx *Context) Path() string {
	return ctx.path
}

func (ctx *Context) SetPath(path string) {
	ctx.path = path
}

//Response
func (ctx *Context) Error(err error) {
	var code int
	var content string

	if herr, ok := err.(*HTTPError); ok {
		code = herr.Code
	} else {
		code = InternalServerError.Code
	}
	content = err.Error()

	ctx.response.SetHeader(HeaderContentType, MIMEApplicationJSONCharsetUTF8)
	ctx.response.SetStatus(code)

	body := map[string]string{"error": content}

	if err := json.NewEncoder(ctx.response.ResponseWriter).Encode(body); err != nil {
		log.Error(err)
	}
}

func (ctx *Context) Redirect(code int, url string) {
	if code < 300 || code > 308 {
		log.Error(errors.New("InvalidRedirectError"))
	}

	ctx.response.SetHeader(HeaderLocation, url)
	ctx.response.SetStatus(code)

	// No 3xx body on POST and PUT
	if ctx.Method() == "GET" {
		note := "<a href=\"" + url + "\">Redirect</a>.\n"
		ctx.response.SetBody([]byte(note))
	}
}

func (ctx *Context) SetCookie(cookie *http.Cookie) {
	if v := cookie.String(); v != "" {
		ctx.response.AddHeader(HeaderSetCookie, v)
	}
}

func (ctx *Context) XML(code int, body interface{}) {
	ctx.response.SetHeader(HeaderContentType, MIMEApplicationXMLCharsetUTF8)
	ctx.response.SetStatus(code)

	if err := json.NewEncoder(ctx.response.ResponseWriter).Encode(body); err != nil {
		log.Error(err)
	}
}

func (ctx *Context) JSON(code int, body interface{}) {
	ctx.response.SetHeader(HeaderContentType, MIMEApplicationJSONCharsetUTF8)
	ctx.response.SetStatus(code)

	if err := json.NewEncoder(ctx.response.ResponseWriter).Encode(body); err != nil {
		log.Error(err)
	}
}

func (ctx *Context) Text(code int, str string, params ...interface{}) {
	ctx.response.SetHeader(HeaderContentType, MIMETextPlainCharsetUTF8)
	ctx.response.SetStatus(code)

	body := fmt.Sprintf(str, params...)

	if _, err := ctx.response.SetBody([]byte(body)); err != nil {
		ctx.Log.Error(err)
	}
}

func (ctx *Context) HTML(code int, str string, params ...interface{}) {
	ctx.response.SetHeader(HeaderContentType, MIMETextHTMLCharsetUTF8)
	ctx.response.SetStatus(code)

	if _, err := fmt.Fprintf(ctx.response.ResponseWriter, str, params...); err != nil {
		ctx.Log.Error(err)
	}
}

func (ctx *Context) File(file string) (err error) {
	f, err := os.Open(file)
	if err != nil {
		log.Error(NotFoundError)
	}
	defer f.Close()

	fi, _ := f.Stat()
	if fi.IsDir() {
		file = filepath.Join(file, indexPage)
		f, err = os.Open(file)
		if err != nil {
			log.Error(NotFoundError)
		}
		defer f.Close()
		if fi, err = f.Stat(); err != nil {
			log.Error(err)
		}
	}
	// Handle HTTP Range properly
	http.ServeContent(ctx.response.ResponseWriter, ctx.request, fi.Name(), fi.ModTime(), f)
	return
}
