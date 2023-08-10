package model

import (
	"context"

	"golangchain/common"
	"golangchain/generation"
	"golangchain/message"

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

func (opm *OpenAIModel) SetOptions(opts ...common.Options) {
	for _, opt := range opts {
		opt(opm.OpenAIModelOption)
	}
}

func (opm *OpenAIModel) Generate(messages []message.Message) (*generation.Generation, error) {
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
		return nil, err
	}

	m := structs.Map(resp)

	return &generation.Generation{
		Text: resp.Choices[0].Message.Content,
		All:  m,
	}, nil
}

func NewOpenAIModel(apiKey string, opts ...common.Options) *OpenAIModel {
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

func WithMaxToken(maxToken int) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*OpenAIModelOption); ok {
			options.MaxToken = maxToken
		}
	}
}
func WithModelName(modelName string) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*OpenAIModelOption); ok {
			options.ModelName = modelName
		}
	}
}

func WithTemperature(temperature float32) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*OpenAIModelOption); ok {
			options.Temperature = temperature
		}
	}
}

func OpenAITokenCountFunc() common.TokenCountHandler {
	return func(messages []message.Message) int {
		panic("nothing to do")
	}
}
