package bash_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/takumin/gyaml/internal/command/completion/bash"
	"github.com/takumin/gyaml/internal/config"
)

func TestNewCommands(t *testing.T) {
	var stdout, stderr bytes.Buffer
	cmd := bash.NewCommands(config.NewConfig(), []cli.Flag{})
	cmd.Writer = &stdout
	cmd.ErrWriter = &stderr

	if cmd.Name != "bash" {
		t.Errorf("expected command name to be 'bash', but got '%s'", cmd.Name)
	}

	if cmd.Usage != "bash completion" {
		t.Errorf("expected command usage to be 'bash completion', but got '%s'", cmd.Usage)
	}

	if err := cmd.Run(context.Background(), []string{}); err != nil {
		t.Errorf("falied to run bash completion: %v\nstdout: %v\n stderr: %v", err, stdout, stderr)
	}

	if stdout.Len() == 0 {
		t.Error("expected bash completion output to be non-empty, but got empty")
	}
}
