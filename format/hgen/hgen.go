package hgen

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
