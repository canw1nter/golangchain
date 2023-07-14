package prompts

import (
	"bytes"
	"fmt"
	"text/template"

	"golangchain/utils"
)

type PromptOption struct {
	Variables map[string]interface{}
}

type Prompt struct {
	Template string
	*PromptOption
}

func (p *Prompt) SetOptions(opts ...utils.Options) {
	for _, opt := range opts {
		opt(p.PromptOption)
	}
}

func (p *Prompt) GetText() string {
	var w = bytes.NewBuffer(nil)

	t, err := template.New("123").Parse(p.Template)
	if err != nil {
		// todo handle err
		fmt.Println(err.Error())
		return ""
	}

	err = t.Execute(w, p.Variables)
	if err != nil {
		// todo handle err
		fmt.Println(err.Error())
		return ""
	}

	return string(w.Bytes())
}

func NewPrompt(template string, opts ...utils.Options) *Prompt {
	prompt := &Prompt{
		Template: template,
		PromptOption: &PromptOption{
			Variables: nil,
		},
	}

	prompt.SetOptions(opts...)

	return prompt
}

func WithVariables(variables map[string]interface{}) utils.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*PromptOption); ok {
			options.Variables = variables
		}
	}
}
