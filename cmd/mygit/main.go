package main

import (
	"fmt"
	"github.com/codecrafters-io/git-starter-go/pkg/commands"
	"os"
)

var Commands = map[string]func() error{
	"init":     commands.Init,
	"cat-file": commands.CatFile,
}

func main() {
	if len(os.Args) < 2 {
		exit("usage: mygit <command> [<args>...]")
	}

	command := os.Args[1]
	if err := runCommand(command); err != nil {
		exit(err.Error())
	}
}

func exit(message string) {
	_, _ = fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func runCommand(command string) error {
	if fn, ok := Commands[command]; ok {
		return fn()
	}
	return fmt.Errorf("unknown command %s", command)
}
