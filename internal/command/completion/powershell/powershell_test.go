package powershell_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/takumin/gyaml/internal/command/completion/powershell"
	"github.com/takumin/gyaml/internal/config"
)

func TestNewCommands(t *testing.T) {
	var stdout, stderr bytes.Buffer
	cmd := powershell.NewCommands(config.NewConfig(), []cli.Flag{})
	cmd.Writer = &stdout
	cmd.ErrWriter = &stderr

	if cmd.Name != "powershell" {
		t.Errorf("expected command name to be 'powershell', but got '%s'", cmd.Name)
	}

	if cmd.Usage != "powershell completion" {
		t.Errorf("expected command usage to be 'powershell completion', but got '%s'", cmd.Usage)
	}

	if err := cmd.Run(context.Background(), []string{}); err != nil {
		t.Errorf("falied to run powershell completion: %v\nstdout: %v\n stderr: %v", err, stdout, stderr)
	}

	if stdout.Len() == 0 {
		t.Error("expected powershell completion output to be non-empty, but got empty")
	}
}
