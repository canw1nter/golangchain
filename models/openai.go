package models

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type OpenAIModel struct {
	APIKey string
	*OpenAIModelOption
}

type OpenAIModelOption struct {
	MaxToken    int
	ModelName   string
	Temperature float32
}

type OptionSet func(option *OpenAIModelOption)

func (opm *OpenAIModel) Generate() {
	client := openai.NewClient(opm.APIKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       opm.ModelName,
			MaxTokens:   opm.MaxToken,
			Temperature: opm.Temperature,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func NewOpenAIModel(apiKey string, optionSets ...OptionSet) *OpenAIModel {
	option := &OpenAIModelOption{
		ModelName:   "gpt-3.5-turbo-0613",
		MaxToken:    400,
		Temperature: 0.0,
	}

	for _, optionSet := range optionSets {
		optionSet(option)
	}

	model := &OpenAIModel{
		APIKey: apiKey,
	}
	model.OpenAIModelOption = option

	return model
}
