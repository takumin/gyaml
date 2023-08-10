package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/takumin/gyaml/internal/command/completion"
	"github.com/takumin/gyaml/internal/command/validation"
	"github.com/takumin/gyaml/internal/config"
)

var (
	AppName  string = "gyaml"
	AppDesc  string = "Golang YAML Tool"
	Version  string = "unknown"
	Revision string = "unknown"
)

func main() {
	cfg := config.NewConfig(
		config.Directory("."),
		config.Includes("yml,yaml"),
	)

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "log-level",
			Aliases:     []string{"l"},
			Usage:       "log level",
			EnvVars:     []string{"LOG_LEVEL"},
			Value:       cfg.LogLevel,
			Destination: &cfg.LogLevel,
		},
	}

	cmds := []*cli.Command{
		completion.NewCommands(cfg, flags),
		validation.NewCommands(cfg, flags),
	}

	app := &cli.App{
		Name:                 AppName,
		Usage:                AppDesc,
		Version:              fmt.Sprintf("%s (%s)", Version, Revision),
		Flags:                flags,
		Commands:             cmds,
		EnableBashCompletion: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
