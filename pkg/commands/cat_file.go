package commands

import (
	"github.com/codecrafters-io/git-starter-go/pkg/git"
	"os"
)

func CatFile() error {
	// TODO: make arg parser

	// TODO: switch on parameters e.g. -p should call smth like git.PrettyPrintHash
	return git.CatFile(git.CatFileOptions{
		ShowSHAContent: &os.Args[3],
	})
}
