package bash

import (
	"context"
	"strings"
	"text/template"

	"github.com/urfave/cli/v3"

	"github.com/takumin/gyaml/internal/config"
)

const bashCompletion = `
#!/bin/bash

_cli_bash_autocomplete() {
	if [[ "${COMP_WORDS[0]}" != "source" ]]; then
		local cur opts base
		COMPREPLY=()
		cur="${COMP_WORDS[COMP_CWORD]}"
		if [[ "$cur" == "-"* ]]; then
			opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --generate-shell-completion )
		else
			opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-shell-completion )
		fi
		COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
		return 0
	fi
}

complete -o bashdefault -o default -o nospace -F _cli_bash_autocomplete {{.}}
`

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:     "bash",
		Usage:    "bash completion",
		HideHelp: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			t := template.Must(template.New("bashCompletion").Parse(strings.TrimSpace(bashCompletion) + "\n"))
			return t.Execute(cmd.Writer, cmd.Name)
		},
	}
}
