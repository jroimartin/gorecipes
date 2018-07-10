package one

import (
	"fmt"

	"github.com/jroimartin/gorecipes/cli/internal/base"
)

var CmdOne = &base.Command{
	UsageLine: "one [-f]",
	Short:     "short description for one",
	Long: `
Long description for one.

It supports multiple lines.
	`,
}

var exampleFlag bool

func init() {
	CmdOne.Run = runOne
	CmdOne.Flag.BoolVar(&exampleFlag, "f", false, "Example flag")
}

func runOne(cmd *base.Command, args []string) {
	fmt.Printf("one! exampleFlag=%v\n", exampleFlag)
}
