package main

import (
	"github.com/alecthomas/kong"
)

var (
	cli struct {
		File string `arg:"" required:"" type:"existingfile" help:"Input file"`
		Output string `short:"o" default:"a.out" help:"Output executable file"`
		Assembly bool `short:"s" help:"Output assembly code"`
	}
)

func main() {
	ctx := kong.Parse(&cli)

	code, err := Parse(cli.File)
	ctx.FatalIfErrorf(err)

	asm, err := Assembly(code, cli.File)
	ctx.FatalIfErrorf(err)

	err = Compile(asm, cli.Output, cli.Assembly)
	ctx.FatalIfErrorf(err)
}
