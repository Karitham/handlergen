package openapi

import (
	"io"
	"regexp"
	"strings"

	"github.com/Karitham/handlergen/gen"
	"gopkg.in/yaml.v2"
)

type Oapi struct {
	Paths map[string]map[string]Route `yaml:"paths"`
}

type Route struct {
	operationID string `yaml:"operationId"`
	RequestBody struct {
		Content  map[string]DataType `json:"content"`
		Required bool                `yaml:"required"`
	} `yaml:"requestBody"`
	Parameters []Parameter `yaml:"parameters"`
}

type DataType struct {
	Schema struct {
		Ref string `json:"$ref"`
	} `json:"schema"`
}

type Parameter struct {
	Schema Schema `yaml:"schema"`
	In     string `yaml:"in"`
	Name   string `yaml:"name"`
}

type Schema struct {
	Type string `yaml:"type"`
}

func Parse(r io.Reader) (gen.Template, error) {
	s := &Oapi{}
	err := yaml.NewDecoder(r).Decode(s)
	if err != nil {
		return gen.Template{}, err
	}
	t := gen.Template{
		Imports: []gen.Import{
			{Path: "net/http"},
		},
	}

	for p, path := range s.Paths {
		for k, route := range path {
			gf := gen.Function{
				Name: formatOAPIName(k, p, route.operationID),
			}

			if route.RequestBody.Required {
				for _, b := range route.RequestBody.Content {
					b.Schema.Ref = strings.TrimPrefix(b.Schema.Ref, "#/definitions/")
				}
				gf.Body = gen.Body{
					Type: "json.RawMessage",
				}
				t.Imports = append(t.Imports, gen.Import{Path: "encoding/json"})
			}

			for _, param := range route.Parameters {
				parsers[param.In](&t, &gf, param)
			}

			gf.HasBody = gf.Body.Type != ""
			gf.HasQueryParams = len(gf.QueryParams) > 0
			if !gf.HasBody && !gf.HasQueryParams {
				continue
			}

			t.Functions = append(t.Functions, gf)
		}
	}

	return t, err
}

var sanitizeRegex = regexp.MustCompile(`\{([\w\d]+)\}`)

func formatOAPIName(op, path, name string) string {
	if name != "" {
		return strings.Title(name)
	}
	path = sanitizeRegex.ReplaceAllString(path, "By/$1")
	paths := strings.Split(strings.Title(path), "/")

	new := strings.Builder{}
	new.WriteString(strings.Title(op))
	for _, p := range paths {
		new.WriteString(strings.Title(p))
	}

	return new.String()
}

func parseTypesQuery(t *gen.Template, typ string) string {
	t.Imports = append(t.Imports, gen.Import{Path: imports[typ]})

	switch typ {
	case "int", "integer":
		return "int"
	case "array":
		return "[]string"
	case "string":
		return "string"
	}
	return typ
}

func parseTypesPath(t *gen.Template, typ string) string {
	t.Imports = append(t.Imports, gen.Import{Path: "github.com/go-chi/chi/v5"})
	t.Imports = append(t.Imports, gen.Import{Path: imports[typ]})

	switch typ {
	case "int", "integer":
		return "int_query"
	case "array":
		return "[]string_query"
	case "string":
		return "string_query"
	}
	return typ
}

func parseTypesHeader(t *gen.Template, typ string) string {
	t.Imports = append(t.Imports, gen.Import{Path: imports[typ]})
	switch typ {
	case "int", "integer":
		return "int_header"
	case "array":
		return "[]string_header"
	case "string":
		return "string_header"
	}
	return typ
}
