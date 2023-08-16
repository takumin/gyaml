package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

var re = regexp.MustCompile(`^line (\d+): (.*)$`)

type ParseError struct {
	Path    string
	Line    int
	Column  int
	Message string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("failed to %s:%d:%d %s", e.Path, e.Line, e.Column, e.Message)
}

func Parse(path string, data []byte) ([]*ParseError, error) {
	var errs []*ParseError
	var obj interface{}

	if err := yaml.Unmarshal(data, &obj); err != nil {
		if te, ok := err.(*yaml.TypeError); ok {
			for _, e := range te.Errors {
				if s := re.FindStringSubmatch(e); len(s) > 1 {
					line, err := strconv.Atoi(s[1])
					if err != nil {
						return nil, err
					}
					errs = append(errs, &ParseError{
						Path:    path,
						Line:    line,
						Message: s[2],
					})
				} else {
					return nil, fmt.Errorf("unknown yaml type error: '%s'", e)
				}
			}
		} else {
			errs = append(errs, &ParseError{
				Path:    path,
				Message: strings.TrimPrefix(err.Error(), "yaml: "),
			})
		}
	}

	if len(errs) > 0 {
		return errs, nil
	} else {
		return nil, nil
	}
}
