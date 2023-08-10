package validation

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/gyaml/internal/config"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "directory",
			Aliases:     []string{"d"},
			Usage:       "search base directory",
			EnvVars:     []string{"DIRECTORY"},
			Value:       cfg.Path.Directory,
			Destination: &cfg.Path.Directory,
		},
		&cli.StringFlag{
			Name:        "includes",
			Aliases:     []string{"i"},
			Usage:       "include files extensions",
			EnvVars:     []string{"INCLUDES"},
			Value:       cfg.Extention.Includes,
			Destination: &cfg.Extention.Includes,
		},
		&cli.StringFlag{
			Name:        "excludes",
			Aliases:     []string{"e"},
			Usage:       "exclude files extensions",
			EnvVars:     []string{"EXCLUDES"},
			Value:       cfg.Extention.Excludes,
			Destination: &cfg.Extention.Excludes,
		},
	}...)
	return &cli.Command{
		Name:    "validation",
		Aliases: []string{"validate", "valid", "v"},
		Usage:   "yaml file validation",
		Flags:   flags,
		Action:  action(cfg),
	}
}

func action(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return nil
	}
}
