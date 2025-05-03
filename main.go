package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/takumin/gyaml/internal/command/completion"
	"github.com/takumin/gyaml/internal/command/validation"
	"github.com/takumin/gyaml/internal/config"
	"github.com/takumin/gyaml/internal/metadata"
	"github.com/takumin/gyaml/internal/version"
)

func main() {
	cfg := config.NewConfig(
		config.Path("."),
		config.Include("**/*.yml"),
		config.Include("**/*.yaml"),
		config.ReportType("rdjsonl"),
	)

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "log-level",
			Aliases:     []string{"l"},
			Usage:       "log level",
			Sources:     cli.EnvVars("LOG_LEVEL"),
			Value:       cfg.LogLevel,
			Destination: &cfg.LogLevel,
		},
	}

	cmds := []*cli.Command{
		completion.NewCommands(cfg, flags),
		validation.NewCommands(cfg, flags),
	}

	app := &cli.Command{
		Name:                  metadata.AppName(),
		Usage:                 metadata.AppDesc(),
		Version:               fmt.Sprintf("%s (%s)", version.Version(), version.Revision()),
		Flags:                 flags,
		Commands:              cmds,
		EnableShellCompletion: true,
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
