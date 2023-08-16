package validation

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"github.com/takumin/gyaml/internal/config"
	"github.com/takumin/gyaml/internal/filelist"
	"github.com/takumin/gyaml/internal/helpers"
	"github.com/takumin/gyaml/internal/parser"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.MultiStringFlag{
			Target: &cli.StringSliceFlag{
				Name:    "include",
				Aliases: []string{"i"},
				Usage:   "include file extension",
				EnvVars: []string{"INCLUDE"},
			},
			Value:       cfg.Extention.Includes,
			Destination: &cfg.Extention.Includes,
		},
		&cli.MultiStringFlag{
			Target: &cli.StringSliceFlag{
				Name:    "exclude",
				Aliases: []string{"e"},
				Usage:   "exclude file extension",
				EnvVars: []string{"EXCLUDE"},
			},
			Value:       cfg.Extention.Excludes,
			Destination: &cfg.Extention.Excludes,
		},
	}...)
	return &cli.Command{
		Name:            "validation",
		Aliases:         []string{"validate", "valid", "v"},
		Usage:           "yaml file validation",
		ArgsUsage:       "[file or directory...]",
		HideHelpCommand: true,
		Flags:           flags,
		Before:          before(cfg),
		Action:          action(cfg),
	}
}

func before(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		if ctx.NArg() >= 1 {
			cfg.Paths = helpers.RemoveDuplicateStrings(ctx.Args().Slice())
		}
		return nil
	}
}

func action(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		paths := make([]string, 0, 1024)

		for _, v := range cfg.Paths {
			info, err := os.Stat(v)
			if err != nil {
				return err
			}

			if info.IsDir() {
				temp, err := filelist.Filelist(
					os.DirFS(v),
					v,
					cfg.Extention.Includes,
					cfg.Extention.Excludes,
				)
				if err != nil {
					return err
				}
				paths = append(paths, temp...)
			} else {
				paths = append(paths, v)
			}
		}

		for _, path := range paths {
			data, err := os.ReadFile(filepath.Clean(path))
			if err != nil {
				return err
			}

			errs, err := parser.Parse(path, data)
			if err != nil {
				return err
			}
			// TODO: report rdjsonl
			for _, e := range errs {
				fmt.Println(e)
			}
		}

		return nil
	}
}
