package gen

import (
	"bytes"
	"embed"
	"io"
	"os"
	"text/template"

	"github.com/davecgh/go-spew/spew"
)

//go:embed templates/*
var tmpl embed.FS

func Execute(t Template, w io.Writer) error {
	tpl, err := template.New("main.gotmpl").ParseFS(tmpl, "templates/*.gotmpl")
	if err != nil {
		return err
	}

	spew.Fdump(os.Stderr, t)

	buf := &bytes.Buffer{}
	if err := tpl.ExecuteTemplate(buf, "main.gotmpl", &t); err != nil {
		return err
	}

	return Format(buf, w)
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
