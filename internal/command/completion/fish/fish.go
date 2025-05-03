package fish

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/takumin/gyaml/internal/config"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:     "fish",
		Usage:    "fish completion",
		HideHelp: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fish, err := cmd.ToFishCompletion()
			if err != nil {
				return err
			}
			if _, err := fmt.Fprint(cmd.Writer, fish); err != nil {
				return err
			}
			return nil
		},
	}
}
