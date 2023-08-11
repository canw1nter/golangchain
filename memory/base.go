package memory

import (
	"golangchain/common"
	"golangchain/memory/history"
	"golangchain/message"

	"github.com/pkg/errors"
)

type IMemory interface {
	GetMemory() ([]message.Message, error)
	ClearMemory()
	SaveToMemory([]message.Message)
}

type MemoryOption struct {
	TokenCount common.TokenCountHandler
	TokenLimit int
	History    history.IHistory
}

type Memory struct {
	messages []message.Message
	*MemoryOption
}

func (m *Memory) SetOptions(opts ...common.Options) {
	for _, opt := range opts {
		opt(m.MemoryOption)
	}
}

func (m *Memory) GetMemory() ([]message.Message, error) {
	if len(m.messages) == 0 && m.History != nil {
		histories, err := m.History.Get()
		if err != nil {
			return nil, errors.Wrap(err, "get memory failed")
		}
		m.messages = *histories
	}

	if m.TokenLimit > 0 {
		for m.TokenCount(m.messages) > m.TokenLimit {
			m.messages = m.messages[:len(m.messages)-1]
		}
	}

	return m.messages, nil
}

func (m *Memory) ClearMemory() {
	m.messages = make([]message.Message, 0)
	if m.History != nil {
		m.History.Clear()
	}
}

func (m *Memory) SaveToMemory(messages []message.Message) {
	if m.History != nil {
		m.History.Add(messages)
	}

	m.messages = append(messages, m.messages...)

	if m.TokenLimit > 0 {
		for m.TokenCount(m.messages) > m.TokenLimit {
			m.messages = m.messages[:len(m.messages)-1]
		}
	}
}

func NewMemory(opts ...common.Options) *Memory {
	memory := &Memory{
		messages: make([]message.Message, 0),
		MemoryOption: &MemoryOption{
			History:    nil,
			TokenLimit: 0,
			TokenCount: nil,
		},
	}

	memory.SetOptions(opts...)

	return memory
}

func WithTokenLimit(tokenLimit int, tokenCount common.TokenCountHandler) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*MemoryOption); ok {
			options.TokenLimit = tokenLimit
			options.TokenCount = tokenCount
		}
	}
}

func WithHistory(history history.IHistory) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*MemoryOption); ok {
			options.History = history
		}
	}
}
