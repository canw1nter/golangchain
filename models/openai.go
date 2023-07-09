package models

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type OpenAIModel struct {
	ModelName   string
	Temperature float32
	APIKey      string
	MaxToken    int
}

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
