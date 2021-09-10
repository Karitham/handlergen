package gen

import (
	"io"

	"mvdan.cc/gofumpt/format"
)

func Format(in io.Reader, out io.Writer) error {
	b, err := io.ReadAll(in)
	if err != nil {
		return err
	}

	res, err := format.Source(b, format.Options{
		LangVersion: "1.16",
		ExtraRules:  true,
	})
	if err != nil {
		return err
	}

	_, err = out.Write(res)
	return err
}
