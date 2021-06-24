package format

import "github.com/Karitham/handlergen/gen"

func mapBasic(f *Structure, pkg string) gen.Template {
	t := gen.Template{
		Imports: []gen.Import{
			{Path: "net/http"},
		},
		PkgName: pkg,
	}

	for name, fun := range f.Functions {
		if fun.Body.Type != "" {
			t.Imports = appendOnceImports(t.Imports, gen.Import{Path: "encoding/json"})
		}

		t.Functions = append(t.Functions, gen.Function{
			Body:           gen.Body{Type: fun.Body.Type},
			Name:           name,
			QueryParams:    queryParamsFromQuery(&t, fun),
			HasQueryParams: len(fun.Query) > 0,
			HasBody:        fun.Body.Type != "",
		})
	}

	return t
}

func queryParamsFromQuery(t *gen.Template, f Function) []gen.QueryParam {
	q := []gen.QueryParam{}

	for n, p := range f.Query {
		q = append(q, gen.QueryParam{
			Name: n,
			Type: parseTypesQuery(t, p.Type),
		})
		if p.Import != "" {
			t.Imports = append(t.Imports, gen.Import{Path: f.Body.Import})
		}
	}
	for n, p := range f.Header {
		q = append(q, gen.QueryParam{
			Name: n,
			Type: parseTypesHeader(t, p.Type),
		})
		if p.Import != "" {
			t.Imports = append(t.Imports, gen.Import{Path: f.Body.Import})
		}
	}
	for n, p := range f.Path {
		q = append(q, gen.QueryParam{
			Name: n,
			Type: parseTypesPath(t, p.Type),
		})
		if p.Import != "" {
			t.Imports = append(t.Imports, gen.Import{Path: f.Body.Import})
		}
	}
	return q
}
