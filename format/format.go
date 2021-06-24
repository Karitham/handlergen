package format

import (
	"io"

	"github.com/Karitham/handlergen/gen"
	"gopkg.in/yaml.v3"
)

func Parse(r io.Reader, format string, pkg string) (gen.Template, error) {
	t := gen.Template{}
	switch format {
	case "", "handlergen":
		s := &Structure{}
		err := yaml.NewDecoder(r).Decode(s)
		if err != nil {
			return gen.Template{}, err
		}
		t = mapBasic(s, pkg)
	case "openapi":
		s := &Oapi{}
		err := yaml.NewDecoder(r).Decode(s)
		if err != nil {
			return gen.Template{}, err
		}
		t = mapOapi(s, pkg)
	}

	return t, nil
}

// Structure defines the default structure of a document
type Structure struct {
	Functions map[string]Function `yaml:"functions"`
}

// Function
type Function struct {
	Query  map[string]Fields `yaml:"query"`
	Path   map[string]Fields `yaml:"path"`
	Header map[string]Fields `yaml:"header"`
	Body   Fields            `yaml:"body"`
}

// Fields
type Fields struct {
	Type   string `yaml:"type"`
	Import string `yaml:"import"`
}
