package validate_test

import (
	"testing"

	"github.com/takumin/gyaml/internal/validate"
)

func TestValidateSuccess(t *testing.T) {
	path := "valid.yaml"
	data := []byte("a: a")
	if _, err := validate.Validate(path, data); err != nil {
		t.Errorf("expected validate error to be 'nil', but got '%s'", err)
	}
}

func TestValidateYAMLError(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		data   []byte
		errors []struct {
			line    int
			column  int
			message string
		}
	}{
		{
			name: "YAML Unknown Error",
			path: "invalid.yaml",
			data: []byte(">:"),
			errors: []struct {
				line    int
				column  int
				message string
			}{
				{
					line:    0,
					column:  0,
					message: `did not find expected comment or line break`,
				},
			},
		},
		{
			name: "YAML Type Error - Duplicate Keys",
			path: "duplicate_keys.yaml",
			data: []byte("a: a\na: a"),
			errors: []struct {
				line    int
				column  int
				message string
			}{
				{
					line:    2,
					column:  0,
					message: `mapping key "a" already defined at line 1`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs, err := validate.Validate(tt.path, tt.data)
			if err != nil {
				t.Errorf("expected error to be 'nil', but got '%s'", err)
			}

			if len(tt.errors) != len(errs) {
				t.Errorf("expected validate errors length to be '%d', but got '%d'", len(tt.errors), len(errs))
			}

			for i, e := range errs {
				if len(e.Error()) <= 0 {
					t.Errorf("expected validate error %d result to be not empty", i)
				}
				if e.Path() != tt.path {
					t.Errorf("expected validate error %d path to be '%s', but got '%s'", i, tt.path, e.Path())
				}
				if e.Line() != tt.errors[i].line {
					t.Errorf("expected validate error %d line to be '%d', but got '%d'", i, tt.errors[i].line, e.Line())
				}
				if e.Column() != tt.errors[i].column {
					t.Errorf("expected validate error %d column to be '%d', but got '%d'", i, tt.errors[i].column, e.Column())
				}
				if e.Message() != tt.errors[i].message {
					t.Errorf("expected validate error %d message to be '%s', but got '%s'", i, tt.errors[i].message, e.Message())
				}
			}
		})
	}
}
