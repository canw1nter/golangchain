package memory

import (
	"golangchain/common"
	"golangchain/memory/history"
	"golangchain/message"
)

type IMemory interface {
	GetMemory() []message.Message
	ClearMemory()
	SaveToMemory([]message.Message)
}

type MemoryOption struct {
	TokenCount common.TokenCountHandler
	TokenLimit int
}

type Memory struct {
	messages []message.Message
	buffered bool
	History  history.IHistory
	*MemoryOption
}

func (m *Memory) SetOptions(opts ...common.Options) {
	for _, opt := range opts {
		opt(m.MemoryOption)
	}
}

func (m *Memory) GetMemory() (*[]message.Message, error) {
	if len(m.messages) == 0 {
		histories, err := m.History.Get()
		if err != nil {
			return nil, err
		}
		m.messages = *histories
	}

	if m.TokenLimit > 0 {
		for m.TokenCount(m.messages) > m.TokenLimit {
			m.messages = m.messages[:len(m.messages)-1]
		}
	}

	return &m.messages, nil
}

func (m *Memory) ClearMemory() {
	m.messages = make([]message.Message, 0)
	m.History.Clear()
}

func (m *Memory) SaveToMemory(messages []message.Message) {
	m.History.Add(messages)
	m.messages = append(messages, m.messages...)

	if m.TokenLimit > 0 {
		for m.TokenCount(m.messages) > m.TokenLimit {
			m.messages = m.messages[:len(m.messages)-1]
		}
	}
}
