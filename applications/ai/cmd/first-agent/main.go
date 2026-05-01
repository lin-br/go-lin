package main

import (
	"log"

	"github.com/lin-br/go-lin/applications/ai/configs"
)

const propertiesFilePath = "configs.yaml"

func main() {
	properties, err := configs.LoadConfigs(propertiesFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if properties.OpenAIKey == nil {
		log.Fatal("OpenAI API key is missing in configuration")
	}

	log.Println(properties.OpenAIKey)
}
