package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

var re = regexp.MustCompile(`^line (\d+): (.*)$`)

type ParseError interface {
	Error() string
	Path() string
	Line() int
	Column() int
	Message() string
}

type parseError struct {
	path    string
	line    int
	column  int
	message string
}

func (e *parseError) Error() string {
	return fmt.Sprintf("failed to %s:%d:%d %s", e.path, e.line, e.column, e.message)
}

func (e *parseError) Path() string {
	return e.path
}

func (e *parseError) Line() int {
	return e.line
}

func (e *parseError) Column() int {
	return e.column
}

func (e *parseError) Message() string {
	return e.message
}

func Parse(path string, data []byte) ([]ParseError, error) {
	var errs []*parseError
	var obj interface{}
	if err := yaml.Unmarshal(data, &obj); err != nil {
		if te, ok := err.(*yaml.TypeError); ok {
			for _, e := range te.Errors {
				if s := re.FindStringSubmatch(e); len(s) > 1 {
					line, err := strconv.Atoi(s[1])
					if err != nil {
						return nil, err
					}
					errs = append(errs, &parseError{
						path:    path,
						line:    line,
						message: s[2],
					})
				} else {
					return nil, fmt.Errorf("unknown yaml type error: '%s'", e)
				}
			}
		} else {
			errs = append(errs, &parseError{
				path:    path,
				message: strings.TrimPrefix(err.Error(), "yaml: "),
			})
		}
	}

	if len(errs) > 0 {
		res := make([]ParseError, 0, len(errs))
		for _, v := range errs {
			res = append(res, v)
		}
		return res, nil
	} else {
		return nil, nil
	}
}
