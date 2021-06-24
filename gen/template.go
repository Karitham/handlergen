package gen

import (
	"bytes"
	"embed"
	"io"
	"strings"
	"text/template"
)

//go:embed templates/*
var tmpl embed.FS

func Execute(t Template, w io.Writer) error {
	templ := template.New("main.gotmpl")
	templ.Funcs(template.FuncMap{"StripTypeSuffix": StripTypeSuffix})

	tpl, err := templ.ParseFS(tmpl, "templates/*.gotmpl")
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := tpl.ExecuteTemplate(buf, "main.gotmpl", &t); err != nil {
		return err
	}

	if err := Format(buf, w); err != nil {
		return err
	}
	return nil
}

func StripTypeSuffix(s string) string {
	return strings.SplitN(s, "_", 2)[0]
}

type Template struct {
	PkgName   string
	Functions []Function
	Imports   []Import
}

type Import struct {
	Path string
}

type Function struct {
	Body           Body
	Name           string
	QueryParams    []QueryParam
	HasQueryParams bool
	HasBody        bool
}

type QueryParam struct {
	Name string
	Type string
}

type Body struct {
	Type string
}
