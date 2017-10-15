package vidar

import (
	"net/http"
)

var Parsers = map[string]Parser{}

type Parser interface {
	ParserName() string
	Parse(obj interface{}, req *http.Request) error
}

func NewParser(name string) Parser {
	if p, ok := Parsers[name]; ok {
		return p
	}

	return new(DefaultParser)
}

func AddParser(name string, p Parser) {
	Parsers[name] = p
}

type DefaultParser struct{}

func (dp *DefaultParser) ParserName() string {
	return "DefaultParser"
}

func (dp *DefaultParser) Parse(obj interface{}, req *http.Request) error {
	return NotImplementedError
}
