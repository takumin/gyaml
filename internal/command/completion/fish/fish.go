package fish

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/takumin/gyaml/internal/config"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:     "fish",
		Usage:    "fish completion",
		HideHelp: true,
		Action: func(ctx *cli.Context) error {
			fish, err := ctx.App.ToFishCompletion()
			if err != nil {
				return err
			}
			if _, err := fmt.Fprint(ctx.App.Writer, fish); err != nil {
				return err
			}
			return nil
		},
	}
}
