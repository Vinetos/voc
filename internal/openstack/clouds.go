package openstack

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type CloudConfig struct {
	Clouds map[string]interface{} `yaml:"clouds"`
}

func ListClouds(configPath string) ([]string, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read cloud.yaml: %w", err)
	}

	var config CloudConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse cloud.yaml: %w", err)
	}

	var cloudNames []string
	for name := range config.Clouds {
		cloudNames = append(cloudNames, name)
	}

	return cloudNames, nil
}
