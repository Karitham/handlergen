package hgen

import "github.com/Karitham/handlergen/gen"

var imports = map[string]string{
	"int":     "strconv",
	"integer": "strconv",
	"array":   "strings",
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
