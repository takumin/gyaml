package zsh

import (
	"context"
	"strings"
	"text/template"

	"github.com/urfave/cli/v3"

	"github.com/takumin/gyaml/internal/config"
)

const zshCompletion = `
#compdef {{.}}

_cli_zsh_autocomplete() {
	local -a opts
	local cur

	cur=${words[-1]}
	if [[ "$cur" == "-"* ]]; then
		opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} ${cur} --generate-shell-completion)}")
	else
		opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-shell-completion)}")
	fi

	if [[ "${opts[1]}" != "" ]]; then
		_describe 'values' opts
	else
		_files
	fi

	return
}

compdef _cli_zsh_autocomplete {{.}}
`

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:     "zsh",
		Usage:    "zsh completion",
		HideHelp: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			t := template.Must(template.New("zshCompletion").Parse(strings.TrimSpace(zshCompletion) + "\n"))
			return t.Execute(cmd.Writer, cmd.Name)
		},
	}
}
