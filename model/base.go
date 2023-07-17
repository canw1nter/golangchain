package model

import (
	"golangchain/generation"
	"golangchain/message"
)

type ILanguageModel interface {
	Generate(messages []message.Message) *generation.Generation
}
