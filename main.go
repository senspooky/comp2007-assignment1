package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/senspooky/comp2007-assignment1/abstraction"
	"github.com/senspooky/comp2007-assignment1/demos"
	exceptions "github.com/senspooky/comp2007-assignment1/exception"
	"github.com/senspooky/comp2007-assignment1/expressivity"
	"github.com/senspooky/comp2007-assignment1/simplicityorthogonality"
	"github.com/senspooky/comp2007-assignment1/syntax"
	"github.com/senspooky/comp2007-assignment1/typechecking"
	"github.com/senspooky/comp2007-assignment1/types"
)

const (
	demoFlag = "demos"
)

var (
	demoFlags arrayFlags
)

type arrayFlags []string

// String is an implementation of the flag.Value interface
func (i *arrayFlags) String() string {
	return fmt.Sprintf("%v", *i)
}

// Set is an implementation of the flag.Value interface
func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTION]...\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprint(flag.CommandLine.Output(), " Examples:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "   %s -%s %s %s\n", os.Args[0], demoFlag, "simplicity", "orthogonality")
		fmt.Fprintf(flag.CommandLine.Output(), "   %s %s %s\n", os.Args[0], "simplicity", "orthogonality")
	}
	flag.Var(&demoFlags, demoFlag, fmt.Sprintf("`Names of the demos to run` which is one or more of the following, seperated by a space: simplicity, orthogonality, datatypes, syntaxdesign, abstraction, expressivity, typechecking, exceptionhandling or restrictedaliasing. The \"-%s\" flag is optional.", demoFlag))
	flag.Parse()
	// default
	for _, arg := range flag.Args() {
		demoFlags = append(demoFlags, arg)
	}

	if len(demoFlags) == 0 {
		flag.Usage()
		return
	}

	toRun := make([]demos.Demo, 0, len(flag.Args()))

	for _, arg := range demoFlags {
		var demo demos.Demo

		switch arg {
		case "simplicityorthogonality":
			demo = simplicityorthogonality.SimplicityDemo()
		case "datatypes":
			demo = types.DataTypesDemo()
		case "syntax":
			demo = syntax.SyyntaxDemo()
		case "abstraction":
			demo = abstraction.AbstractionDemo()
		case "expressivity":
			demo = expressivity.ExpressivityDemo()
		case "typechecking":
			demo = typechecking.TypeCheckingDemo()
		case "exceptionhandling":
			demo = exceptions.ExceptionDemo()
		default:
			flag.Usage()
			return
		}

		toRun = append(toRun, demo)
	}

	for _, demo := range toRun {
		demo.Demo()
	}
}
