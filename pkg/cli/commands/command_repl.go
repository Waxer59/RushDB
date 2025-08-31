package commands

import (
	"bufio"
	"context"
	"os"

	"github.com/Waxer59/RushDB/internal/utils"
	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
)

func SetUpRepl() *cli.Command {
	return &cli.Command{
		Name:  "repl",
		Usage: "Starts a REPL",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return startRepl()
		},
	}
}

func startRepl() error {
	for {
		c := color.New(color.FgHiBlack).Add(color.Bold)

		if _, err := c.Print("RushDB > "); err != nil {
			return err
		}

		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()

		text := scanner.Text()

		switch text {
		case "exit":
			return nil
		case "clear", "cls":
			utils.CallClearConsoleSc()
			continue
		}
	}
}
