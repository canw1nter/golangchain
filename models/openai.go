package models

import (
	"context"
	"fmt"
	"github.com/fatih/structs"
	"github.com/sashabaranov/go-openai"
	"golangchain/generations"
	"golangchain/messages"
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

func (opm *OpenAIModel) Generate(messages []messages.Message) *generations.Generation {
	client := openai.NewClient(opm.APIKey)

	var messagesParam []openai.ChatCompletionMessage
	for _, m := range messages {
		openaiM := openai.ChatCompletionMessage{
			Role:    m.Role,
			Content: m.Content,
		}

		messagesParam = append(messagesParam, openaiM)
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       opm.ModelName,
			MaxTokens:   opm.MaxToken,
			Temperature: opm.Temperature,
			Messages:    messagesParam,
		},
	)

	if err != nil {
		fmt.Printf("err: %v", err)
		return nil
	}

	m := structs.Map(resp)

	return &generations.Generation{
		Text: resp.Choices[0].Message.Content,
		All:  m,
	}
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
