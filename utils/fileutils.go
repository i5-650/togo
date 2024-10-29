package utils

import (
	"gopkg.in/yaml.v3"
	"os"
)

// LoadYAML reads and parses a YAML file into the provided struct
func LoadYAML(filename string, out interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, out)
}
