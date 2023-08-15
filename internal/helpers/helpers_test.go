package helpers_test

import (
	"reflect"
	"testing"

	"github.com/takumin/gyaml/internal/helpers"
)

func TestRemoveDuplicateStrings(t *testing.T) {
	data := []string{".", "."}
	want := []string{"."}
	got := helpers.RemoveDuplicateStrings(data)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected string slice to be '%s', but got '%s'", want, got)
	}
}
