package prompt

import (
	"bytes"
	"text/template"

	"golangchain/common"

	"github.com/pkg/errors"
)

type PromptOption struct {
	Variables map[string]interface{}
}

type Prompt struct {
	Template string
	*PromptOption
}

func (p *Prompt) SetOptions(opts ...common.Options) {
	for _, opt := range opts {
		opt(p.PromptOption)
	}
}

func (p *Prompt) GetText() (string, error) {
	var w = bytes.NewBuffer(nil)

	t, err := template.New("").Parse(p.Template)
	if err != nil {
		return "", errors.Wrap(err, "get prompt text failed")
	}

	err = t.Execute(w, p.Variables)
	if err != nil {
		return "", errors.Wrap(err, "get prompt text failed")
	}

	return string(w.Bytes()), nil
}

func NewPrompt(template string, opts ...common.Options) *Prompt {
	prompt := &Prompt{
		Template: template,
		PromptOption: &PromptOption{
			Variables: nil,
		},
	}

	prompt.SetOptions(opts...)

	return prompt
}

func WithVariables(variables map[string]interface{}) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*PromptOption); ok {
			options.Variables = variables
		}
	}
}
