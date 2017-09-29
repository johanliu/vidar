package parser

import (
	"net/http"
	"strings"

	"github.com/johanliu/vidar/constant"
	"github.com/ugorji/go/codec"
)

type MsgpackParser struct{}

func (mp *MsgpackParser) PluginName() string {
	return "MsgpackParser"
}

func (mp *MsgpackParser) Parse(obj interface{}, req *http.Request) error {
	if !strings.HasPrefix(req.Header.Get(constant.HeaderContentType), constant.MIMEApplicationMsgpack) {
		return constant.UnsupportedMediaTypeError
	}

	if err := codec.NewDecoder(req.Body, new(codec.MsgpackHandle)).Decode(obj); err != nil {
		return constant.BadRequestError
	}

	return nil
}
