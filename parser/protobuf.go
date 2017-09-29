package parser

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/johanliu/vidar/constant"
)

type ProtobufParser struct{}

func (pp *ProtobufParser) PluginName() string {
	return "JSONParser"
}

func (pp *ProtobufParser) Parse(obj interface{}, req *http.Request) error {
	if !strings.HasPrefix(req.Header.Get(constant.HeaderContentType), constant.MIMEApplicationProtobuf) {
		return constant.UnsupportedMediaTypeError
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return constant.BadRequestError
	}

	if err = proto.Unmarshal(buf, obj.(proto.Message)); err != nil {
		return constant.BadRequestError
	}

	return nil
}
