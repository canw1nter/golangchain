package message

import "fmt"

type IMessage interface {
	ToString() string
}

type Message struct {
	Role    string
	Content string
}

func (m *Message) ToString() string {
	return fmt.Sprintf("%s:%s", m.Role, m.Content)
}
