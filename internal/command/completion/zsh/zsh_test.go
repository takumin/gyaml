package zsh_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/takumin/gyaml/internal/command/completion/zsh"
	"github.com/takumin/gyaml/internal/config"
)

func TestNewCommands(t *testing.T) {
	var stdout, stderr bytes.Buffer
	cmd := zsh.NewCommands(config.NewConfig(), []cli.Flag{})
	cmd.Writer = &stdout
	cmd.ErrWriter = &stderr

	if cmd.Name != "zsh" {
		t.Errorf("expected command name to be 'zsh', but got '%s'", cmd.Name)
	}

	if cmd.Usage != "zsh completion" {
		t.Errorf("expected command usage to be 'zsh completion', but got '%s'", cmd.Usage)
	}

	if err := cmd.Run(context.Background(), []string{}); err != nil {
		t.Errorf("falied to run zsh completion: %v\nstdout: %v\n stderr: %v", err, stdout, stderr)
	}

	if stdout.Len() == 0 {
		t.Error("expected zsh completion output to be non-empty, but got empty")
	}
}
