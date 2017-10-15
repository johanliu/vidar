package json

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/johanliu/vidar"
)

type JSONParser struct{}

func (jp *JSONParser) ParserName() string {
	return "JSONParser"
}

func (jp *JSONParser) Parse(obj interface{}, req *http.Request) error {
	if !strings.HasPrefix(req.Header.Get(vidar.HeaderContentType), vidar.MIMEApplicationJSON) {
		return vidar.UnsupportedMediaTypeError
	}

	if err := json.NewDecoder(req.Body).Decode(obj); err != nil {
		return vidar.BadRequestError
	}

	return nil
}

func init() {
	vidar.AddParser("JSON", &JSONParser{})
}
