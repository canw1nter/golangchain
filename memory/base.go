package memory

import (
	"golangchain/memory/history"
	"golangchain/message"
)

type Memory struct {
	History history.IHistory
}

type IMemory interface {
	GetMemory() []message.Message
	ClearMemory()
	SaveToMemory([]message.Message)
}
