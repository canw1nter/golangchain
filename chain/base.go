package chain

import (
	"golangchain/common"
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

func (c *Chain) SetOptions(opts ...common.Options) {
	for _, opt := range opts {
		opt(c.ChainOption)
	}
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

func NewChain(opts ...common.Options) (*Chain, error) {
	chain := &Chain{}

	chain.SetOptions(opts...)

	return chain, nil
}

func WithMemory(memory memory.IMemory) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*ChainOption); ok {
			options.Memory = memory
		}
	}
}

func WithModel(model model.ILanguageModel) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*ChainOption); ok {
			options.Model = model
		}
	}
}

func WithPrompt(prompt prompt.Prompt) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*ChainOption); ok {
			options.Prompt = prompt
		}
	}
}

func WithNextChain(nextChain IChain) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*ChainOption); ok {
			options.NextChain = nextChain
		}
	}
}
