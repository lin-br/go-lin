package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lin-br/go-lin/applications/ai/configs"
	"github.com/tmc/langchaingo/llms"

	"github.com/tmc/langchaingo/llms/openai"
)

const propertiesFilePath = "configs.yaml"
const modelName = "gpt-5.4-mini"

func main() {
	properties := getConfigs()
	llm, err := openai.New(openai.WithModel(modelName), openai.WithToken(*properties.OpenAIKey))

	if err != nil {
		log.Fatal(err)
	}

	prompt := "write a haiku about ai"
	completion, err := llms.GenerateFromSinglePrompt(
		context.Background(),
		llm,
		prompt,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)
}

func getConfigs() *configs.Config {
	properties, err := configs.LoadConfigs(propertiesFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if properties.OpenAIKey == nil {
		log.Fatal("OpenAI API key is missing in configuration")
	}
	return properties
}
