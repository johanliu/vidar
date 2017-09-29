package vidar

import (
	"net/http"

	"github.com/johanliu/vidar/constant"
	"github.com/johanliu/vidar/parser"
)

type ParserPlugins interface {
	PluginName() string
	Parse(obj interface{}, req *http.Request) error
}

func (ctx *Context) NewParser(name string) ParserPlugins {
	switch name {
	case "JSON":
		return new(parser.JSONParser)
	case "XML":
		return new(parser.XMLParser)
	case "Protobuf":
		return new(parser.ProtobufParser)
	case "Msgpack":
		return new(parser.MsgpackParser)
	default:
		return new(DefaultParser)
	}
}

type DefaultParser struct{}

func (dp *DefaultParser) PluginName() string {
	return "DefaultParser"
}

func (dp *DefaultParser) Parse(obj interface{}, req *http.Request) error {
	return constant.NotImplementedError
}
