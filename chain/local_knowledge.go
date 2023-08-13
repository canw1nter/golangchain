package chain

import (
	"github.com/pkg/errors"
	"golangchain/common"
	"golangchain/prompt"
	"strings"
)

const localKnowledgePromptTemplate = `Analyze Human's true intentions based on history messages and Human's last question, and derive helpful answers to Human based on context
history messages: 
{{.History}}

human last question: {{.Question}}

context: 
{{.Context}}
`

type LocalKnowledgeChain struct {
	*Chain
}

func (lkc *LocalKnowledgeChain) Run(inputs map[string]interface{}) (interface{}, error) {
	if !lkc.verifyInputKeys(inputs) {
		return nil, errors.New("local knowledge chain run failed, inputs are incorrect")
	}

	histories, err := lkc.Memory.GetMemory()
	if err != nil {
		return nil, errors.Wrap(err, "local knowledge chain run failed")
	}
	historiesStr := make([]string, len(histories))
	for _, history := range histories {
		historiesStr = append(historiesStr, history.ToString())
	}
	inputs["History"] = strings.Join(historiesStr, "\n")

	return nil, nil
}

func NewLocalKnowledgeChain(opts ...common.Options) *LocalKnowledgeChain {
	lkc := &LocalKnowledgeChain{
		Chain: &Chain{
			ChainOption: &ChainOption{
				Prompt:     *prompt.NewPrompt(localKnowledgePromptTemplate),
				InputKeys:  []string{"Question"},
				OutputKeys: []string{"Answer"},
			},
		},
	}

	lkc.SetOptions(opts...)

	return lkc
}
