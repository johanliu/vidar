package parser

import (
	"encoding/xml"
	"net/http"

	"github.com/johanliu/vidar/constant"
)

type XMLParser struct{}

func (xp *XMLParser) PluginName() string {
	return "XMLParser"
}

func (xp *XMLParser) Parse(obj interface{}, req *http.Request) error {
	if err := xml.NewDecoder(req.Body).Decode(obj); err != nil {
		return constant.BadRequestError
	}

	return nil
}
