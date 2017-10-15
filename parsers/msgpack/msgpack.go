package parsers

import (
	"net/http"
	"strings"

	"github.com/johanliu/vidar"
	"github.com/ugorji/go/codec"
)

type MsgpackParser struct{}

func (mp *MsgpackParser) ParserName() string {
	return "MsgpackParser"
}

func (mp *MsgpackParser) Parse(obj interface{}, req *http.Request) error {
	if !strings.HasPrefix(req.Header.Get(vidar.HeaderContentType), vidar.MIMEApplicationMsgpack) {
		return vidar.UnsupportedMediaTypeError
	}

	if err := codec.NewDecoder(req.Body, new(codec.MsgpackHandle)).Decode(obj); err != nil {
		return vidar.BadRequestError
	}

	return nil
}

func init() {
	vidar.AddParser("Msgpack", &JSONParser{})
}
