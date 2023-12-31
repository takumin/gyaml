package config_test

import (
	"reflect"
	"testing"

	"github.com/takumin/gyaml/internal/config"
)

func TestLogLevel(t *testing.T) {
	want := &config.Config{LogLevel: "TEST"}
	got := &config.Config{}
	config.LogLevel("TEST").Apply(got)
	if !reflect.DeepEqual(want, got) {
		t.Error("expected config struct to be equal, but got not equal")
	}
}

func TestPath(t *testing.T) {
	want := &config.Config{}
	want.Paths = []string{"TEST"}
	got := &config.Config{}
	config.Path("TEST").Apply(got)
	if !reflect.DeepEqual(want, got) {
		t.Error("expected config struct to be equal, but got not equal")
	}
}

func TestIncludes(t *testing.T) {
	want := &config.Config{}
	want.Extention.Includes = []string{"TEST"}
	got := &config.Config{}
	config.Include("TEST").Apply(got)
	if !reflect.DeepEqual(want, got) {
		t.Error("expected config struct to be equal, but got not equal")
	}
}

func TestExcludes(t *testing.T) {
	want := &config.Config{}
	want.Extention.Excludes = []string{"TEST"}
	got := &config.Config{}
	config.Exclude("TEST").Apply(got)
	if !reflect.DeepEqual(want, got) {
		t.Error("expected config struct to be equal, but got not equal")
	}
}

func TestReportType(t *testing.T) {
	want := &config.Config{Report: struct{ Type string }{Type: "TEST"}}
	got := &config.Config{}
	config.ReportType("TEST").Apply(got)
	if !reflect.DeepEqual(want, got) {
		t.Error("expected config struct to be equal, but got not equal")
	}
}
