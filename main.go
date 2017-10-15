package main

import (
	"fmt"
	"os"
)

func main() {
	conf, ok := NewConfig(os.Args[1:], os.Stderr)
	if !ok {
		os.Exit(1)
	}

	cli, err := NewCLI(conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = cli.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cli.Close()
}
