package main

import (
	"github.com/alecthomas/kong"
)

var (
	cli struct {
		File string `arg:"" required:"" type:"existingfile" help:"Input file"`
		Output string `short:"o" default:"a.out" help:"Name of the executable file."`
		Include string `short:"I" default:"." type:"existingdir" help:"Path to the MPI header file."`
		Lib string `short:"L" default:"." type:"existingdir" help:"Path to the MPI library."`
		Assembly bool `short:"S" help:"Whether to output assembly code."`
		Optimize int `short:"O" enum:"0,1" default:"1" help:"Optimization level."`
		Verbose bool `short:"v" help:"Whether to print verbose information."`
	}
)

func main() {
	ctx := kong.Parse(&cli)

	code, hasMpi, err := Parse(cli.File)
	ctx.FatalIfErrorf(err)

	if cli.Optimize > 0 {
		code = Optimize(code)
	}

	s, err := Assembly(code, cli.File)
	ctx.FatalIfErrorf(err)

	c := Library(hasMpi)

	err = CompileAndLink(s, c, cli.Output, cli.Include, cli.Lib, hasMpi, cli.Assembly, cli.Verbose)
	ctx.FatalIfErrorf(err)
}
