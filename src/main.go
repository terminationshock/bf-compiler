package main

import (
	"github.com/alecthomas/kong"
)

var (
	cli struct {
		File string `arg:"" required:"" type:"existingfile" help:"Input file"`
		Output string `short:"o" default:"a.out" help:"Name of the executable file."`
		Assembly bool `short:"S" help:"Whether to output assembly code."`
		Optimize int `short:"O" enum:"0,1" default:"1" help:"Optimization level."`
		StackSize int `default:"30000" help:"Stack size."`
		Verbose bool `short:"v" help:"Whether to print verbose information."`
	}
)

func main() {
	ctx := kong.Parse(&cli)

	code, err := Parse(cli.File)
	ctx.FatalIfErrorf(err)

	if cli.Optimize > 0 {
		code = Optimize(code)
	}

	assembly, err := Assembly(code, cli.File, cli.StackSize, cli.Optimize > 0)
	ctx.FatalIfErrorf(err)

	err = CompileAndLink(assembly, cli.Output, cli.Assembly, cli.Verbose)
	ctx.FatalIfErrorf(err)
}
