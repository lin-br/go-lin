package configs

import (
	"os"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	OpenAIKey *string `yaml:"openai_api_key"`
}

func LoadConfigs(filePath string) (*Config, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	expandedContent := os.ExpandEnv(string(file))

	properties := &Config{}

	if err := yaml.Unmarshal([]byte(expandedContent), properties); err != nil {
		return nil, err
	}
	return properties, nil
}
