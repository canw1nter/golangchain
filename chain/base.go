package chain

import (
	"golangchain/common"
	"golangchain/memory"
	"golangchain/model"
	"golangchain/prompt"
)

type IChain interface {
	Run(map[string]interface{}) (map[string]interface{}, error)
	verifyInputKeys(inputs map[string]interface{}) bool
}

type ChainOption struct {
	NextChain  IChain
	Memory     memory.IMemory
	Model      model.ILanguageModel
	Prompt     prompt.Prompt
	InputKeys  []string
	OutputKeys []string
}

type Chain struct {
	*ChainOption
}

func (c *Chain) SetOptions(opts ...common.Options) {
	for _, opt := range opts {
		opt(c.ChainOption)
	}
}

func (c *Chain) verifyInputKeys(inputs map[string]interface{}) bool {
	for _, inputKey := range c.InputKeys {
		if _, ok := inputs[inputKey]; !ok {
			return false
		}
	}
	return true
}

func (c *Chain) Run(inputs map[string]interface{}) (map[string]interface{}, error) {
	// do something
	// then run next chain if not nil
	if c.NextChain != nil {
		return c.NextChain.Run(inputs)
	} else {
		return inputs, nil
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

func WithInputKeys(inputKeys []string) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*ChainOption); ok {
			options.InputKeys = inputKeys
		}
	}
}

func WithOutputKeys(outputKeys []string) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*ChainOption); ok {
			options.OutputKeys = outputKeys
		}
	}
}
