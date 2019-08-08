package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config contains general app settings
type Config struct {
	Port    string `yaml:"port"`
	Storage string `yaml:"storage"`
}

// LoadYamlConfig loads config from yaml file to struct
func LoadYamlConfig(path string) (Config, error) {
	var config Config
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	return config, err
}
