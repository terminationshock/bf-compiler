package main

import (
	"github.com/alecthomas/kong"
)

var (
	cli struct {
		File string `arg:"" required:"" type:"existingfile" help:"Input file"`
		Output string `short:"o" default:"a.out" help:"Name of the executable file."`
		Lib string `short:"L" default:"." type:"existingdir" help:"Path to the MPI library."`
		Include string `short:"I" default:"." type:"existingdir" help:"Path to the MPI header file."`
		Assembly bool `short:"s" help:"Whether to output assembly code."`
		Verbose bool `short:"v" help:"Whether to print verbose information."`
	}
)

func main() {
	ctx := kong.Parse(&cli)

	code, hasMpi, err := Parse(cli.File)
	ctx.FatalIfErrorf(err)

	asm, err := Assembly(code, cli.File)
	ctx.FatalIfErrorf(err)

	c := Library(hasMpi)

	err = Compile(asm, c, cli.Output, cli.Include, cli.Lib, hasMpi, cli.Assembly, cli.Verbose)
	ctx.FatalIfErrorf(err)
}
