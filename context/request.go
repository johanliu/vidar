package context

import (
	"errors"
	"mime/multipart"
)

type parsePlugins interface {
}

const MaxMemory = 64 << 20 // 64MB

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

func (ctx *Context) Form(key string, defaultvalues ...string) string {
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

/*
func (ctx *Context) getRaw() interface{} {

}*/

func (param *Parameters) Params(key string) (string, bool) {
	return param.getParams(key)
}

func (params *Parameters) getParams(key string) (string, bool) {
	for _, param := range *params {
		if param.key == key {
			return param.value, true
		}
	}

	return "", false
}
