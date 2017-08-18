package parser

import (
	"encoding/xml"
	"net/http"

	"github.com/johanliu/Vidar/constant"
)

type XMLParser struct{}

func (jp *XMLParser) PluginName() string {
	return "XMLParser"
}

func (jp *XMLParser) Parse(obj interface{}, req *http.Request) error {
	if err := xml.NewDecoder(req.Body).Decode(obj); err != nil {
		return constant.BadRequestError
	}

	return nil
}
