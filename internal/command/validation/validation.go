package validation

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"

	"github.com/urfave/cli/v3"

	"github.com/takumin/gyaml/internal/config"
	"github.com/takumin/gyaml/internal/filelist"
	"github.com/takumin/gyaml/internal/parser"
	"github.com/takumin/gyaml/internal/report"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "type",
			Aliases:     []string{"t"},
			Usage:       "report type",
			Sources:     cli.EnvVars("TYPE"),
			Value:       cfg.Report.Type,
			Destination: &cfg.Report.Type,
		},
		&cli.StringSliceFlag{
			Name:        "include",
			Aliases:     []string{"i"},
			Usage:       "include file extension",
			Sources:     cli.EnvVars("INCLUDE"),
			Value:       cfg.Extention.Includes,
			Destination: &cfg.Extention.Includes,
		},
		&cli.StringSliceFlag{
			Name:        "exclude",
			Aliases:     []string{"e"},
			Usage:       "exclude file extension",
			Sources:     cli.EnvVars("EXCLUDE"),
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

func before(cfg *config.Config) func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
	return func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
		if cmd.NArg() >= 1 {
			s := cmd.Args().Slice()
			sort.Strings(s)
			cfg.Paths = slices.Compact(s)
		}
		return ctx, nil
	}
}

func action(cfg *config.Config) func(ctx context.Context, cmd *cli.Command) error {
	return func(ctx context.Context, cmd *cli.Command) error {
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

		errs := make(map[string][]*parser.ParseError, len(paths))
		for _, path := range paths {
			data, err := os.ReadFile(filepath.Clean(path))
			if err != nil {
				return err
			}
			perrs, err := parser.Parse(path, data)
			if err != nil {
				return err
			}
			if perrs != nil {
				errs[path] = perrs
			}
		}

		keys := make([]string, 0, len(errs))
		for k := range errs {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			for _, e := range errs[k] {
				switch cfg.Report.Type {
				case "rdjsonl":
					buf, err := report.ReviewdogDiagnosticJSONLines(e.Path, e.Message, e.Line, e.Column)
					if err != nil {
						return err
					}
					if _, err := fmt.Fprintln(cmd.Writer, string(buf)); err != nil {
						return err
					}
				default:
					return fmt.Errorf("unsupported report type: %s", cfg.Report.Type)
				}
			}
		}

		return nil
	}
}
