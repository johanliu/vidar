package context

import "github.com/johanliu/Vidar/logger"

func (ctx *Context) Form(key string, defaultvalues ...string) []string {
	value, ok := ctx.getForm(key)
	if !ok {
		if len(defaultvalues) > 0 {
			return defaultvalues
		}
	}
	return value
}

func (ctx *Context) getForm(key string) ([]string, bool) {
	if err := ctx.Request.ParseForm(); err != nil {
		logger.Error.Println("Parse HTTP Form failed")
	}

	values, exists := ctx.Request.Form[key]
	if exists {
		return values, true
	}

	return nil, false
}

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
