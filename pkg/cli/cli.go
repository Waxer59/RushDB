package cli

import (
	"context"
	"os"

	"github.com/Waxer59/RushDB/pkg/cli/commands"
	"github.com/urfave/cli/v3"
)

func SetUp() {
	(&cli.Command{
		Name:        "RushDB",
		Version:     "0.0.1",
		Usage:       "A CLI tool for managing RushDB",
		Description: "A CLI tool for managing RushDB",
		Commands: []*cli.Command{
			commands.SetUpRepl(),
		},
	}).Run(context.Background(), os.Args)
}
