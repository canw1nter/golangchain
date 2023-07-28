package memory

import (
	"golangchain/memory/history"
	"golangchain/message"
)

type IMemory interface {
	GetMemory() []message.Message
	ClearMemory()
	SaveToMemory([]message.Message)
}

type Memory struct {
	History    history.IHistory
	messages   []message.Message
	tokenCount func(messages []message.Message) int
	TokenLimit int
}

func (m *Memory) GetMemory() ([]message.Message, error) {
	if m.messages == nil || len(m.messages) == 0 {
		histories, err := m.History.Get()
		if err != nil {
			return nil, err
		}
		m.messages = histories
	}

	if m.TokenLimit > 0 {
		for m.tokenCount(m.messages) > m.TokenLimit {
			m.messages = m.messages[:len(m.messages)-1]
		}
	}

	return m.messages, nil
}

func (m *Memory) ClearMemory() {
	m.messages = make([]message.Message, 0)
	m.History.Clear()
}

func (m *Memory) SaveToMemory(messages []message.Message) {
	m.History.Add(messages)
	m.messages = append(messages, m.messages...)

	if m.TokenLimit > 0 {
		for m.tokenCount(m.messages) > m.TokenLimit {
			m.messages = m.messages[:len(m.messages)-1]
		}
	}
}
