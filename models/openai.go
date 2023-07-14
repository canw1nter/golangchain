package models

import (
	"context"
	"fmt"

	"golangchain/generations"
	"golangchain/messages"
	"golangchain/utils"

	"github.com/fatih/structs"
	"github.com/sashabaranov/go-openai"
)

type OpenAIModelOption struct {
	MaxToken    int
	ModelName   string
	Temperature float32
}

type OpenAIModel struct {
	APIKey string
	*OpenAIModelOption
}

func (opm *OpenAIModel) SetOptions(opts ...utils.Options) {
	for _, opt := range opts {
		opt(opm.OpenAIModelOption)
	}
}

func NewOpenAIModel(apiKey string, opts ...utils.Options) *OpenAIModel {
	model := &OpenAIModel{
		APIKey: apiKey,
		OpenAIModelOption: &OpenAIModelOption{
			ModelName:   "gpt-3.5-turbo-0613",
			MaxToken:    400,
			Temperature: 0.7,
		},
	}

	model.SetOptions(opts...)

	return model
}

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
