package main

import (
	"flag"
	"fmt"
	"github.com/accentdesign/gtml/cmd/gtml/generate"
	"io"
	"log/slog"
	"os"
)

func main() {
	code := run(os.Stdin, os.Stdout, os.Stderr, os.Args)
	if code != 0 {
		os.Exit(code)
	}
}

const usageText = `usage: gtml <command> [<args>...]

templ - build HTML with Go

See https://github.com/accentdesign/gtml for more information.

commands:
  generate   Generates Go code from an html file
  help       Print this help message

`

func run(stdin io.Reader, stdout, stderr io.Writer, args []string) (code int) {
	if len(args) < 2 {
		fmt.Fprint(stderr, usageText)
		return 64 // EX_USAGE
	}
	switch args[1] {
	case "generate":
		return generateCmd(stdin, stdout, stderr, args[2:])
	case "help", "-help", "--help", "-h":
		fmt.Fprint(stdout, usageText)
		return 0
	}
	fmt.Fprint(stderr, usageText)
	return 64 // EX_USAGE
}

func newLogger(stderr io.Writer) *slog.Logger {
	return slog.New(slog.NewTextHandler(stderr, nil))
}

const generateUsageText = `usage: gtml generate [<args>...]

Generates Go code from an html file.

Args:
  -help
    Print this help message.

Examples:

  Generate code for a specific html file to an output file:

    gtml generate < file.html > file.go

  Generate code for a specific html file to stdout:

    gtml generate < file.html

`

func generateCmd(stdin io.Reader, stdout, stderr io.Writer, args []string) (code int) {
	cmd := flag.NewFlagSet("generate", flag.ExitOnError)
	helpFlag := cmd.Bool("help", false, "")
	err := cmd.Parse(args)
	if err != nil {
		fmt.Fprint(stderr, generateUsageText)
		return 64 // EX_USAGE
	}
	if *helpFlag {
		fmt.Fprint(stdout, generateUsageText)
		return
	}

	log := newLogger(stderr)

	if err := generate.Run(log, stdin, stdout); err != nil {
		fmt.Fprintf(stderr, "Command failed: %v\n", err)
		return 1
	}

	return 0
}
