package main

import (
	"flag"
	"os"

	"github.com/Karitham/handlergen/builder"
	"github.com/Karitham/handlergen/gen"
	"github.com/peterbourgon/ff"
	"github.com/rs/zerolog/log"
)

func main() {
	fs := flag.NewFlagSet("handlergen", flag.ExitOnError)
	var (
		filename = fs.String("file", "handlers.yaml", "handlers gen config file")
		pkg      = fs.String("debug", "handlers", "package name")
	)

	err := ff.Parse(fs, os.Args[1:])
	if err != nil {
		log.Fatal().Err(err).Msg("main thread stopped")
	}

	if err := Run(*filename, *pkg); err != nil {
		log.Fatal().Err(err).Msg("main thread stopped")
	}
}

func Run(filename, pkg string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	structure, err := builder.Parse(file)
	if err != nil {
		return err
	}

	if err := gen.Execute(Map(structure, pkg), os.Stdout); err != nil {
		return err
	}

	return nil
}

func Map(f *builder.Structure, pkg string) gen.Template {
	t := gen.Template{
		Imports: []gen.Import{
			{Path: "net/http"},
			{Path: "encoding/json"},
			{Path: "strconv"},
		},
		PkgName: pkg,
	}

	for name, fun := range f.Functions {
		t.Functions = append(t.Functions, gen.Function{
			Body:           gen.Body{Type: fun.Body.Type},
			Name:           name,
			QueryParams:    QueryParamsFromQuery(fun.Query),
			HasQueryParams: len(fun.Query) > 0,
			HasBody:        fun.Body.Type != "",
		})
		t.Imports = append(t.Imports, ImportsFromParams(fun)...)
	}

	return t
}

func QueryParamsFromQuery(fb map[string]builder.Fields) []gen.QueryParam {
	q := []gen.QueryParam{}
	for n, p := range fb {
		q = append(q, gen.QueryParam{
			Name: n,
			Type: p.Type,
		})
	}
	return q
}

func ImportsFromParams(f builder.Function) []gen.Import {
	i := []gen.Import{}
	for _, p := range f.Query {
		if p.Import != "" {
			i = append(i, gen.Import{Path: p.Import})
		}
	}

	if f.Body.Import != "" {
		i = append(i, gen.Import{Path: f.Body.Import})
	}
	return i
}
