package validate

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

var re = regexp.MustCompile(`^line (\d+): (.*)$`)

type ValidaError interface {
	Error() string
	Path() string
	Line() int
	Column() int
	Message() string
}

type validateError struct {
	path    string
	line    int
	column  int
	message string
}

func (e *validateError) Error() string {
	return fmt.Sprintf("failed to %s:%d:%d %s", e.path, e.line, e.column, e.message)
}

func (e *validateError) Path() string {
	return e.path
}

func (e *validateError) Line() int {
	return e.line
}

func (e *validateError) Column() int {
	return e.column
}

func (e *validateError) Message() string {
	return e.message
}

func Validate(path string) ([]ValidaError, error) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	var errs []*validateError
	var obj interface{}
	if err := yaml.Unmarshal(data, &obj); err != nil {
		if te, ok := err.(*yaml.TypeError); ok {
			for _, e := range te.Errors {
				if s := re.FindStringSubmatch(e); len(s) > 1 {
					line, err := strconv.Atoi(s[1])
					if err != nil {
						return nil, err
					}
					errs = append(errs, &validateError{
						path:    path,
						line:    line,
						message: s[2],
					})
				} else {
					return nil, fmt.Errorf("unknown yaml type error: '%s'", e)
				}
			}
		} else {
			errs = append(errs, &validateError{
				path:    path,
				message: strings.TrimPrefix(err.Error(), "yaml: "),
			})
		}
	}

	if len(errs) > 0 {
		res := make([]ValidaError, 0, len(errs))
		for _, v := range errs {
			res = append(res, v)
		}
		return res, nil
	} else {
		return nil, nil
	}
}
