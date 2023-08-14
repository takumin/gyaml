package version_test

import (
	"testing"

	"github.com/takumin/gyaml/internal/version"
)

func TestVersion(t *testing.T) {
	if got := version.Version(); got != "unknown" {
		t.Errorf("expected version to be 'unknown', but got '%s'", got)
	}
}

func TestRevision(t *testing.T) {
	if got := version.Revision(); got != "unknown" {
		t.Errorf("expected revision to be 'unknown', but got '%s'", got)
	}
}
