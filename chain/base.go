package chain

import (
	"golangchain/memory"
	"golangchain/model"
	"golangchain/prompt"
)

type IChain interface {
	Run(inputs map[string]interface{}) (outputs interface{})
}

type ChainOption struct {
	NextChain IChain
	Memory    memory.IMemory
	Model     model.ILanguageModel
	Prompt    prompt.Prompt
}

type Chain struct {
	*ChainOption
}

func (c *Chain) Run(inputs map[string]interface{}) (outputs interface{}) {
	// do something
	// then run next chain if not nil
	if c.NextChain != nil {
		return c.NextChain.Run(inputs)
	} else {
		return outputs
	}
}
