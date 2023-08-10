package chain

import (
	"github.com/pkg/errors"

	"golangchain/common"
	"golangchain/generation"
	"golangchain/message"
	"golangchain/prompt"
)

const promptTemplate = `Summarize the context below as a standalone question

context: {{.Context}}
`

type StandaloneQuestionChain struct {
	*Chain
}

func (sqc *StandaloneQuestionChain) Run(inputs map[string]interface{}) (interface{}, error) {
	if v, ok := inputs["Context"]; ok {
		sqc.Chain.Prompt.Variables["Context"] = v
	} else {
		return nil, errors.New("standalone question chain run failed,not found variable 'Context' in inputs")
	}

	var messages []message.Message
	if sqc.Chain.Memory != nil {
		messages = sqc.Chain.Memory.GetMemory()
	} else {
		messages = make([]message.Message, 1)
	}

	text, err := sqc.Chain.Prompt.GetText()
	if err != nil {
		return nil, errors.Wrap(err, "standalone question chain run failed")
	}
	messages = append(messages, message.Message{Role: "system", Content: text})

	var result *generation.Generation
	if sqc.Chain.Model != nil {
		result, err = sqc.Chain.Model.Generate(messages)
		if err != nil {
			return nil, errors.Wrap(err, "standalone question chain run failed, get generation failed")
		}
	} else {
		return nil, errors.Wrap(err, "standalone question chain run failed, not set model for chain")
	}

	return sqc.Chain.Run(map[string]interface{}{"SDQ": result.Text})
}

func NewStandaloneQuestionChain(opts ...common.Options) *StandaloneQuestionChain {
	sqc := &StandaloneQuestionChain{
		Chain: &Chain{
			ChainOption: &ChainOption{
				Prompt:     *prompt.NewPrompt(promptTemplate),
				inputKeys:  []string{"Context"},
				outputKeys: []string{"SDQ"}, // SDQ - Standalone Question
			},
		},
	}

	sqc.Chain.SetOptions(opts...)

	return sqc
}
