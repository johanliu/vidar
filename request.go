package vidar

import (
	"errors"
	"mime/multipart"
)

type parsePlugins interface {
}

const MaxMemory = 32 << 20 // 64MB

func (ctx *Context) Method() string {
	method := ctx.Request.Method

	return method
}

// Query parameters which pick up k-v pairs from URI queries
// for example: http://localhost:8080/index/department?users=alice
// return {"users":"alice"}
func (ctx *Context) Query(key string, defaultvalues ...string) string {
	value, ok := ctx.getQuery(key)
	if !ok && len(defaultvalues) > 0 {
		return defaultvalues[0]
	}

	return value
}

func (ctx *Context) getQuery(key string) (string, bool) {
	values, exists := ctx.Request.URL.Query()[key]

	if exists && len(values) > 0 {
		return values[0], true
	}

	return "", false
}

// Path parameters which pick up k-v pairs from URI pathes
// for example: http://localhost:8080/index/department/users/alice
// return {"users":"alice"}
// It should be supported by users which defined router for query like
// "/index/department/users/:users"
func (ctx *Context) PathParam(key string, defaultValues ...string) string {
	value, ok := ctx.getPathParam(key)
	if !ok {
		if len(defaultValues) > 0 {
			return defaultValues[0]
		}
	}
	return value
}

func (ctx *Context) getPathParam(key string) (string, bool) {
	//TODO
	return ctx.Parameters.value[2], true
}

// Form paramters which pick k-v pairs from "multipart/form-data" or
// "application/x-www-form-urlencoded"
func (ctx *Context) FormValue(key string, defaultvalues ...string) string {
	value, ok := ctx.getForm(key)
	if !ok {
		if len(defaultvalues) > 0 {
			return defaultvalues[0]
		}
	}
	return value
}

func (ctx *Context) getForm(key string) (string, bool) {
	ctx.Request.ParseMultipartForm(MaxMemory)

	values, exists := ctx.Request.PostForm[key]
	if exists && len(values) > 0 {
		return values[0], true
	}

	values, exists = ctx.Request.MultipartForm.Value[key]
	if exists && len(values) > 0 {
		return values[0], true
	}

	return "", false
}

func (ctx *Context) File(key string) (multipart.File, error) {
	if fh, ok := ctx.getFile(key); ok {
		f, err := fh.Open()
		if err != nil {
			return nil, err
		}
		return f, nil
	}

	return nil, errors.New("File not exists")
}

func (ctx *Context) getFile(key string) (*multipart.FileHeader, bool) {
	ctx.Request.ParseMultipartForm(MaxMemory)

	if files, exists := ctx.Request.MultipartForm.File[key]; exists {
		return files[0], true
	}

	return nil, false
}
