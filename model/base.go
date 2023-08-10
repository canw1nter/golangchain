package model

import (
	"golangchain/common"
	"golangchain/generation"
	"golangchain/message"
)

type ILanguageModel interface {
	Generate(messages []message.Message) (*generation.Generation, error)
	TokenCountFunc() common.TokenCountHandler
}
