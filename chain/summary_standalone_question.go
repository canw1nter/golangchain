package chain

import (
	"golangchain/common"
	"golangchain/prompt"
)

const promptTemplate = `Summarize the context below as a separate question

context: {{.Context}}
`

type StandaloneQuestionChain struct {
	*Chain
}

func NewStandaloneQuestionChain(opts ...common.Options) *StandaloneQuestionChain {
	sqc := &StandaloneQuestionChain{
		Chain: &Chain{
			ChainOption: &ChainOption{
				Prompt: *prompt.NewPrompt(promptTemplate),
			},
		},
	}

	sqc.Chain.SetOptions(opts...)

	return sqc
}
