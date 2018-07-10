package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gorecipes/cli/internal/base"
	"github.com/jroimartin/gorecipes/cli/internal/help"
	"github.com/jroimartin/gorecipes/cli/internal/one"
	"github.com/jroimartin/gorecipes/cli/internal/two"
)

func init() {
	base.Commands = []*base.Command{
		one.CmdOne,
		two.CmdTwo,
	}

	base.Usage = mainUsage
}

func main() {
	log.SetFlags(0)

	flag.Usage = base.Usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		base.Usage()
	}

	if args[0] == "help" {
		help.Help(args[1:])
		return
	}

	for _, cmd := range base.Commands {
		cmd.Flag.Usage = cmd.Usage
		if cmd.Name() == args[0] {
			cmd.Flag.Parse(args[1:])
			args = cmd.Flag.Args()
			cmd.Run(cmd, args)
			return
		}
	}

	fmt.Fprintf(os.Stderr, "cli: unknown subcommand %q\nRun 'cli help' for usage.\n", args[0])
	os.Exit(2)
}

func mainUsage() {
	help.PrintUsage()
	os.Exit(2)
}
