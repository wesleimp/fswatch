package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v2"

	"github.com/wesleimp/fswatch/internal/runner"
)

var errStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:    "fswatch",
		Usage:   "Run commands when file changes",
		Version: "0.1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Usage: "path to watch",
				Value: pwd,
			},
		},
		Action: func(c *cli.Context) error {
			conf := runner.Config{
				Command: c.Args().Slice(),
				Path:    c.String("path"),
			}

			return runner.Run(conf)
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(errStyle.Render(err.Error()))
		os.Exit(1)
	}
}
