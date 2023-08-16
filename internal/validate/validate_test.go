package validate_test

import (
	"testing"

	"github.com/takumin/gyaml/internal/validate"
)

func TestValidateSuccess(t *testing.T) {
	if _, err := validate.Validate("../../testdata/valid/simple.yaml"); err != nil {
		t.Errorf("expected validate error to be 'nil', but got '%s'", err)
	}
}

func TestValidateFaliedReadFile(t *testing.T) {
	if _, err := validate.Validate("noneexist"); err == nil {
		t.Errorf("expected validate error to be not 'nil', but got '%s'", err)
	}
}

func TestValidateYAMLError(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		errors []struct {
			line    int
			column  int
			message string
		}
	}{
		{
			name: "YAML Unknown Error",
			path: "../../testdata/invalid/invalid.yaml",
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
			path: "../../testdata/invalid/duplicate_keys.yaml",
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
				{
					line:    5,
					column:  0,
					message: `mapping key "A" already defined at line 4`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs, err := validate.Validate(tt.path)
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
