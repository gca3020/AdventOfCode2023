package parse

import (
	"os"
	"path/filepath"
	"strings"
)

type Input struct {
	Name string
	Data []byte
}

func ReadInputs() ([]*Input, error) {
	inputs := make([]*Input, 0)
	files, err := filepath.Glob("*.txt")
	if err != nil {
		return nil, err
	}

	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, &Input{
			Name: filepath.Base(path),
			Data: data,
		})
	}
	return inputs, nil
}

func (i *Input) Lines() []string {
	return strings.Split(strings.TrimSpace(string(i.Data)), "\n")
}
