package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/takumin/gyaml/internal/command/completion"
	"github.com/takumin/gyaml/internal/command/validation"
	"github.com/takumin/gyaml/internal/config"
	"github.com/takumin/gyaml/internal/metadata"
	"github.com/takumin/gyaml/internal/version"
)

func main() {
	cfg := config.NewConfig(
		config.Path("."),
		config.Include("yml"),
		config.Include("yaml"),
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
		Name:                 metadata.AppName(),
		Usage:                metadata.AppDesc(),
		Version:              fmt.Sprintf("%s (%s)", version.Version(), version.Revision()),
		Authors:              []*cli.Author{{Name: metadata.AuthorName()}},
		Flags:                flags,
		Commands:             cmds,
		EnableBashCompletion: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
