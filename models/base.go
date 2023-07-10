package models

import (
	"golangchain/generations"
	"golangchain/messages"
)

type IBaseLanguageModel interface {
	Generate(messages []messages.Message) *generations.Generation
}
