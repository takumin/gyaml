package validation

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/takumin/gyaml/internal/config"
	"github.com/takumin/gyaml/internal/filelist"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
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
		Name:      "validation",
		Aliases:   []string{"validate", "valid", "v"},
		Usage:     "yaml file validation",
		ArgsUsage: "[file or directory...]",
		Flags:     flags,
		Before:    before(cfg),
		Action:    action(cfg),
	}
}

func before(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		cfg.Paths = append(cfg.Paths, ctx.Args().Slice()...)
		cfg.Paths = removeDuplicateString(cfg.Paths)
		return nil
	}
}

func action(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		paths := make([]string, 0, 1024)

		for _, v := range cfg.Paths {
			temp, err := filelist.Filelist(
				os.DirFS(v),
				v,
				strings.Split(cfg.Extention.Includes, ","),
				strings.Split(cfg.Extention.Excludes, ","),
			)
			if err != nil {
				return err
			}
			paths = append(paths, temp...)
		}

		// TODO: validation implemented
		for _, path := range paths {
			fmt.Println(path)
		}

		return nil
	}
}

func removeDuplicateString(l []string) (r []string) {
	k := make(map[string]bool)
	for _, s := range l {
		if _, ok := k[s]; !ok {
			k[s] = true
			r = append(r, s)
		}
	}
	return r
}
