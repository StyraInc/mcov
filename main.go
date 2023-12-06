package main

import (
	"fmt"
	"os"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/loader"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mcov <paths>")
		os.Exit(1)
	}

	regos, err := loader.AllRegos(os.Args[1:])
	if err != nil {
		panic(err)
	}

	cmp := ast.NewCompiler()

	cmp.Compile(regos.ParsedModules())

	if version, ok := cmp.Required.MinimumCompatibleVersion(); ok {
		fmt.Println(version)
		os.Exit(0)
	}

	fmt.Println("Could not determine minimum compatible version")
	os.Exit(1)
}
