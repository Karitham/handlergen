package main

import (
	"flag"
	"os"

	"github.com/Karitham/handlergen/format"
	"github.com/Karitham/handlergen/gen"
	"github.com/peterbourgon/ff"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fs := flag.NewFlagSet("handlergen", flag.ExitOnError)
	var (
		file = fs.String("file", "handlers.yaml", "handlers gen config file")
		pkg  = fs.String("pkg", "handlers", "package name")
		fmt  = fs.String("format", "handlergen", "config format, values are handlergen or openapi")
	)

	err := ff.Parse(fs, os.Args[1:])
	if err != nil {
		log.Fatal().Err(err).Msg("main thread stopped")
	}

	if err := Run(*file, *fmt, *pkg); err != nil {
		log.Fatal().Err(err).Msg("main thread stopped")
	}
}

func Run(filename, f, pkg string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	s, err := format.Parse(file, f, pkg)
	if err != nil {
		return err
	}

	if err := gen.Execute(s, os.Stdout); err != nil {
		return err
	}

	return nil
}
