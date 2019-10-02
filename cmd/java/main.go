package main

import (
	"fmt"
	"github.com/mkideal/cli"
)

var (
	version = "0.0.1"
)

type argOp struct {
	cli.Helper
	Version bool   `cli:"v,version" usage:"get jvmgo version"`
	Arg     string `cli:"arg" usage:""`
}

func main() {
	cli.Run(new(argOp), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argOp)
		switch {
		case argv.Version:
			fmt.Println(version)
		default:

		}
		return nil
	})
}