package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Database `yaml:"database"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	file, err := os.Open("./internal/infra/config/properties.yaml")
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func ParseFlags() (string, error) {
	var configPath string
	flag.StringVar(&configPath, "config", "./properties.yaml", "config file path")

	flag.Parse()

	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}
	return configPath, nil
}
