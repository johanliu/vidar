package xml

import (
	"encoding/xml"
	"net/http"

	"github.com/johanliu/vidar"
)

type XMLParser struct{}

func (xp *XMLParser) ParserName() string {
	return "XMLParser"
}

func (xp *XMLParser) Parse(obj interface{}, req *http.Request) error {
	if err := xml.NewDecoder(req.Body).Decode(obj); err != nil {
		return vidar.BadRequestError
	}

	return nil
}

func init() {
	vidar.AddParser("XML", &JSONParser{})
}
