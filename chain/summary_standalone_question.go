package chain

import (
	"github.com/pkg/errors"

	"golangchain/common"
	"golangchain/generation"
	"golangchain/message"
	"golangchain/prompt"
)

const summaryQuestionPromptTemplate = `Summarize the context below as a standalone question

context: {{.Context}}
`

type StandaloneQuestionChain struct {
	*Chain
}

func (sqc *StandaloneQuestionChain) Run(inputs map[string]interface{}) (interface{}, error) {
	if !sqc.verifyInputKeys(inputs) {
		return nil, errors.New("standalone question chain run failed, inputs are incorrect")
	}
	sqc.Prompt.Variables = inputs

	var messages []message.Message
	if sqc.Memory != nil {
		var err error
		messages, err = sqc.Memory.GetMemory()
		if err != nil {
			return nil, errors.Wrap(err, "standalone question chain run failed")
		}
	} else {
		messages = make([]message.Message, 1)
	}

	text, err := sqc.Prompt.GetText()
	if err != nil {
		return nil, errors.Wrap(err, "standalone question chain run failed")
	}
	messages = append(messages, message.Message{Role: "system", Content: text})

	var result *generation.Generation
	if sqc.Model != nil {
		result, err = sqc.Model.Generate(messages)
		if err != nil {
			return nil, errors.Wrap(err, "standalone question chain run failed, get generation failed")
		}
	} else {
		return nil, errors.New("standalone question chain run failed, not set model for chain")
	}

	outputs := make(map[string]interface{})
	for _, outputKey := range sqc.OutputKeys {
		outputs[outputKey] = result.Text
	}

	return sqc.Chain.Run(outputs)
}

func NewStandaloneQuestionChain(opts ...common.Options) *StandaloneQuestionChain {
	sqc := &StandaloneQuestionChain{
		Chain: &Chain{
			ChainOption: &ChainOption{
				Prompt:     *prompt.NewPrompt(summaryQuestionPromptTemplate),
				InputKeys:  []string{"Context"},
				OutputKeys: []string{"SDQ"}, // SDQ - Standalone Question
			},
		},
	}

	sqc.SetOptions(opts...)

	return sqc
}
