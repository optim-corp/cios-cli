package utils

import (
	"gopkg.in/yaml.v2"
)

func LoadYamlStruct(_path string, st interface{}) error {
	byts, err := Path(_path).ReadFile()
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(byts, st)
	if err != nil {
		return err
	}
	return nil
}

func WriteYaml(_path string, yml interface{}) error {
	if b, err := yaml.Marshal(&yml); err != nil {
		return err
	} else {
		Path(_path).WriteFileAsString(string(b))
	}
	return nil
}
