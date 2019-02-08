package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fname := os.Args[1]
	code, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}

	comp := NewCompiler(string(code))
	comp.Compile()
	m := NewMachine(comp.instructions, os.Stdin, os.Stdout)
	m.Execute()
}
