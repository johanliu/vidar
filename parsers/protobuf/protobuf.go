package protobuf

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/johanliu/vidar"
)

type ProtobufParser struct{}

func (pp *ProtobufParser) ParserName() string {
	return "ProtobufParser"
}

func (pp *ProtobufParser) Parse(obj interface{}, req *http.Request) error {
	if !strings.HasPrefix(req.Header.Get(vidar.HeaderContentType), vidar.MIMEApplicationProtobuf) {
		return vidar.UnsupportedMediaTypeError
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return vidar.BadRequestError
	}

	if err = proto.Unmarshal(buf, obj.(proto.Message)); err != nil {
		return vidar.BadRequestError
	}

	return nil
}

func init() {
	vidar.AddParser("Protobuf", &ProtobufParser{})
}
