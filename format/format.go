package format

import (
	"io"

	"github.com/Karitham/handlergen/format/hgen"
	"github.com/Karitham/handlergen/format/openapi"
	"github.com/Karitham/handlergen/gen"
)

var drivers = map[string]func(r io.Reader) (gen.Template, error){
	"openapi":    openapi.Parse,
	"":           hgen.Parse,
	"handlergen": hgen.Parse,
	"hgen":       hgen.Parse,
}

func Parse(r io.Reader, format string, pkg string) (gen.Template, error) {
	t, err := drivers[format](r)
	if err != nil {
		return gen.Template{}, err
	}

	t.Imports = UniqueImports(t.Imports)
	t.PkgName = pkg

	return t, nil
}

// UniqueImports returns a slice of unique imports
func UniqueImports(i []gen.Import) []gen.Import {
	m := map[gen.Import]struct{}{}
	for _, imp := range i {
		if imp.Path == "" {
			continue
		}
		m[imp] = struct{}{}
	}

	// convert map to slice
	var s []gen.Import
	for k := range m {
		s = append(s, k)
	}

	return s
}
