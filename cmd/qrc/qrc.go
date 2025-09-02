package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/mattn/go-colorable"
		isatty "github.com/mattn/go-isatty"
	"io"
	"os"
		"strings"

	"github.com/fumiyas/qrc/lib"
)

var appVersion = "1.0.0" // overridden by -ldflags "-X main.appVersion=<version>"

type cmdOptions struct {
	Help    bool `short:"h" long:"help" description:"show this help message"`
	Inverse bool `short:"i" long:"invert" description:"invert color"`
	Format  string `short:"f" long:"format" choice:"aa" choice:"sixel" description:"output format (aa|sixel); default auto-detect"`
		Version bool `long:"version" description:"print version and exit"`
}

func showHelp() {
	const v = `Usage: qrc [OPTIONS] [TEXT]

Options:
  -h, --help
    Show this help message
  -i, --invert
    Invert color
  -f, --format <aa|sixel>
	Output format (default: auto-detect)
	  aa     ASCII art using ANSI colors
	  sixel  Sixel graphics (if terminal supports it)
      
			If not specified, auto-detect tries to use Sixel only on TTY when
			environment suggests support (QRC_SIXEL=1, TERM contains "sixel" or "mlterm",
			or XTERM_SIXEL=1 for xterm). Otherwise falls back to ASCII art.
	--version
		Print version and exit

Text examples:
  http://www.example.jp/
  MAILTO:foobar@example.jp
  WIFI:S:myssid;T:WPA;P:pass123;;
`

	os.Stderr.Write([]byte(v))
}

func pErr(format string, a ...interface{}) {
	fmt.Fprint(os.Stderr, os.Args[0], ": ")
	fmt.Fprintf(os.Stderr, format, a...)
}

func main() {
	ret := 0
	defer func() { os.Exit(ret) }()

	opts := &cmdOptions{}
	optsParser := flags.NewParser(opts, flags.PrintErrors)
	args, err := optsParser.Parse()
	if err != nil || len(args) > 1 {
		showHelp()
		ret = 1
		return
	}
	if opts.Help {
		showHelp()
		return
	}
	if opts.Version {
		fmt.Fprintf(os.Stdout, "qrc %s\n", appVersion)
		return
	}

	var text string
	if len(args) == 1 {
		text = args[0]
	} else {
		// No args: if stdin is a TTY, show help and exit to avoid hanging
		if isatty.IsTerminal(os.Stdin.Fd()) {
			showHelp()
			ret = 1
			return
		}
		text_bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			pErr("read from stdin failed: %v\n", err)
			ret = 1
			return
		}
		text = string(text_bytes)
	}

	grid, err := qrc.EncodeToGrid(text)
	if err != nil {
		pErr("encode failed: %v\n", err)
		ret = 1
		return
	}

	// Decide output format
	switch opts.Format {
	case "sixel":
		qrc.PrintSixel(os.Stdout, grid, opts.Inverse)
		return
	case "aa":
		stdout := colorable.NewColorableStdout()
		qrc.PrintAA(stdout, grid, opts.Inverse)
		return
	}

	// Heuristic auto-detect: only attempt Sixel on TTY with env hints
	if isatty.IsTerminal(os.Stdout.Fd()) {
		term := os.Getenv("TERM")
		force := os.Getenv("QRC_SIXEL") == "1"
		xtermSixel := os.Getenv("XTERM_SIXEL") == "1"
		if force || strings.Contains(term, "sixel") || strings.Contains(term, "mlterm") || (strings.Contains(term, "xterm") && xtermSixel) {
			qrc.PrintSixel(os.Stdout, grid, opts.Inverse)
			return
		}
	}
	stdout := colorable.NewColorableStdout()
	qrc.PrintAA(stdout, grid, opts.Inverse)
}
