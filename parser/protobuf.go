package parser

import (
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/johanliu/Vidar/constant"
)

type ProtobufParser struct{}

func (pp *ProtobufParser) PluginName() string {
	return "JSONParser"
}

func (pp *ProtobufParser) Parse(obj interface{}, req *http.Request) error {

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return constant.BadRequestError
	}

	if err = proto.Unmarshal(buf, obj.(proto.Message)); err != nil {
		return constant.BadRequestError
	}

	return nil
}
