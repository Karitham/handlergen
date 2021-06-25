package format

import (
	"regexp"
	"strings"

	"github.com/Karitham/handlergen/gen"
)

type Oapi struct {
	Paths map[string]map[string]Route `yaml:"paths"`
}

type Route struct {
	operationID string      `yaml:"operationId"`
	Parameters  []Parameter `yaml:"parameters"`
}

type Parameter struct {
	Schema Schema `yaml:"schema"`
	In     string `yaml:"in"`
	Name   string `yaml:"name"`
}

type Schema struct {
	Type string `yaml:"type"`
}

func mapOapi(f *Oapi, pkg string) gen.Template {
	t := gen.Template{
		Imports: []gen.Import{
			{Path: "net/http"},
		},
		PkgName: pkg,
	}

	for p, path := range f.Paths {
		for k, route := range path {
			gf := gen.Function{
				Name: formatOAPIName(k, p, route.operationID),
			}
			for _, param := range route.Parameters {
				switch param.In {
				case "query":
					gf.QueryParams = append(gf.QueryParams, gen.QueryParam{
						Name: param.Name,
						Type: parseTypesQuery(&t, param.Schema.Type),
					})
				case "path":
					gf.QueryParams = append(gf.QueryParams, gen.QueryParam{
						Name: param.Name,
						Type: parseTypesPath(&t, param.Schema.Type),
					})
				case "header":
					gf.QueryParams = append(gf.QueryParams, gen.QueryParam{
						Name: param.Name,
						Type: parseTypesHeader(&t, param.Schema.Type),
					})
				case "body":
					gf.Body = gen.Body{
						Type: param.Schema.Type,
					}
					t.Imports = appendOnceImports(t.Imports, gen.Import{Path: "encoding/json"})
					if gf.Body.Type != "" {
						t.Imports = appendOnceImports(t.Imports, gen.Import{Path: gf.Body.Type})
					}
				}
			}
			gf.HasBody = gf.Body.Type != ""
			gf.HasQueryParams = len(gf.QueryParams) > 0

			if !gf.HasBody && !gf.HasQueryParams {
				continue
			}

			t.Functions = append(t.Functions, gf)
		}
	}

	return t
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
	switch typ {
	case "int", "integer":
		t.Imports = appendOnceImports(t.Imports, gen.Import{Path: "strconv"})
		return "int"
	case "array":
		t.Imports = appendOnceImports(t.Imports, gen.Import{Path: "strings"})
		return "[]string"
	case "string":
		return "string"
	}
	return typ
}

func parseTypesPath(t *gen.Template, typ string) string {
	t.Imports = appendOnceImports(t.Imports, gen.Import{Path: "github.com/go-chi/chi/v5"})

	switch typ {
	case "int", "integer":
		t.Imports = appendOnceImports(t.Imports, gen.Import{Path: "strconv"})
		return "int_query"
	case "array":
		t.Imports = appendOnceImports(t.Imports, gen.Import{Path: "strings"})
		return "[]string_query"
	case "string":
		return "string_query"
	}
	return typ
}

func parseTypesHeader(t *gen.Template, typ string) string {
	switch typ {
	case "int", "integer":
		t.Imports = appendOnceImports(t.Imports, gen.Import{Path: "strconv"})
		return "int_header"
	case "array":
		t.Imports = appendOnceImports(t.Imports, gen.Import{Path: "strings"})
		return "[]string_header"
	case "string":
		return "string_header"
	}
	return typ
}

func appendOnceImports(imps []gen.Import, i ...gen.Import) []gen.Import {
	for _, s := range imps {
		for _, j := range i {
			if s == j {
				continue
			}
			imps = append(imps, j)
		}
	}
	return imps
}
