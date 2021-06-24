package builder

import (
	"io"

	"gopkg.in/yaml.v3"
)

type Structure struct {
	Functions map[string]Function `yaml:"functions"`
}

func Parse(r io.Reader) (*Structure, error) {
	s := &Structure{}
	return s, yaml.NewDecoder(r).Decode(s)
}

// Function
type Function struct {
	Query map[string]Fields `yaml:"query"`
	Body  Fields            `yaml:"body"`
}

// Fields
type Fields struct {
	Type   string `yaml:"type"`
	Import string `yaml:"import"`
}
