package openapi

import "github.com/Karitham/handlergen/gen"

var imports = map[string]string{
	"int":     "strconv",
	"integer": "strconv",
	"array":   "strings",
}

var parsers = map[string]func(*gen.Template, *gen.Function, Parameter){
	"query": func(t *gen.Template, f *gen.Function, p Parameter) {
		f.QueryParams = append(f.QueryParams, gen.QueryParam{
			Name: p.Name,
			Type: parseTypesQuery(t, p.Schema.Type),
		})
	},
	"path": func(t *gen.Template, f *gen.Function, p Parameter) {
		f.QueryParams = append(f.QueryParams, gen.QueryParam{
			Name: p.Name,
			Type: parseTypesPath(t, p.Schema.Type),
		})
	},
	"header": func(t *gen.Template, f *gen.Function, p Parameter) {
		f.QueryParams = append(f.QueryParams, gen.QueryParam{
			Name: p.Name,
			Type: parseTypesHeader(t, p.Schema.Type),
		})
	},
	"body": func(t *gen.Template, f *gen.Function, p Parameter) {
		f.Body = gen.Body{
			Type: p.Schema.Type,
		}
		t.Imports = append(t.Imports, gen.Import{Path: "encoding/json"})
		if f.Body.Type != "" {
			t.Imports = append(t.Imports, gen.Import{Path: f.Body.Type})
		}
	},
}
