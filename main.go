package main

import (
	"io"
	"os"
)

const (
	ExitOk  = 0
	ExitErr = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	return ExitOk
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	status := cli.Run(os.Args)
	os.Exit(status)
}
