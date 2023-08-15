package validate

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Validate(path string) (string, error) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return "", err
	}

	var obj interface{}
	if err := yaml.Unmarshal(data, &obj); err != nil {
		if te, ok := err.(*yaml.TypeError); ok {
			// TODO: parse error
			for _, v := range te.Errors {
				fmt.Println(v)
			}
		} else {
			// TODO: parse error
			fmt.Println(err)
		}
	}

	// TODO: rdjsonl result
	return "", nil
}
