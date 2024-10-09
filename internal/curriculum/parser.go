package curriculum

import (
	"errors"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func ParseFile(filePath string) (*CV, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	switch path.Ext(filePath) {
	case ".yaml", ".yml":
		return parseYamlFile(buf)
		// TODO case ".json":
	}
	return nil, errors.New(path.Ext(filePath) + " is not a valid file extension")
}

func parseYamlFile(buf []byte) (*CV, error) {
	cv := CV{}
	err := yaml.Unmarshal(buf, &cv)
	if err != nil {
		return nil, err
	}
	return &cv, nil
}
