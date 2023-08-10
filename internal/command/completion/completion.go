package completion

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/gyaml/internal/command/completion/bash"
	"github.com/takumin/gyaml/internal/command/completion/fish"
	"github.com/takumin/gyaml/internal/command/completion/powershell"
	"github.com/takumin/gyaml/internal/command/completion/zsh"
	"github.com/takumin/gyaml/internal/config"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	cmds := []*cli.Command{
		bash.NewCommands(cfg, flags),
		fish.NewCommands(cfg, flags),
		zsh.NewCommands(cfg, flags),
		powershell.NewCommands(cfg, flags),
	}
	return &cli.Command{
		Name:        "completion",
		Usage:       "command completion",
		Subcommands: cmds,
		HideHelp:    true,
	}
}
