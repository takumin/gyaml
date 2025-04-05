package fish_test

import (
	"bytes"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/takumin/gyaml/internal/command/completion/fish"
	"github.com/takumin/gyaml/internal/config"
)

func TestNewCommands(t *testing.T) {
	var stdout, stderr bytes.Buffer
	app := &cli.App{Writer: &stdout, ErrWriter: &stderr}
	app.Setup()
	ctx := cli.NewContext(app, nil, nil)
	cmd := fish.NewCommands(config.NewConfig(), []cli.Flag{})

	if cmd.Name != "fish" {
		t.Errorf("expected command name to be 'fish', but got '%s'", cmd.Name)
	}

	if cmd.Usage != "fish completion" {
		t.Errorf("expected command usage to be 'fish completion', but got '%s'", cmd.Usage)
	}

	if err := cmd.Run(ctx); err != nil {
		t.Errorf("falied to run fish completion: %v\nstdout: %v\n stderr: %v", err, stdout, stderr)
	}

	if stdout.Len() == 0 {
		t.Error("expected fish completion output to be non-empty, but got empty")
	}
}

func TestNewCommandsFailed(t *testing.T) {
	var stdout, stderr bytes.Buffer
	app := &cli.App{Writer: &stdout, ErrWriter: &stderr}
	app.Setup()
	ctx := cli.NewContext(app, nil, nil)
	cmd := fish.NewCommands(config.NewConfig(), []cli.Flag{})

	// Referenced by ToFishCompletion() function.
	cli.FishCompletionTemplate = `{{.}`

	if err := cmd.Run(ctx); err == nil {
		t.Error("expected fish completion result to be failed, but got succeeded")
	}
}
