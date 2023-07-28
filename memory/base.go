package memory

import (
	"golangchain/memory/history"
	"golangchain/message"
)

type Memory struct {
	History history.IHistory
}

type IMemory interface {
	GetAllMemory() []message.Message
	ClearMemory()
	SaveToMemory([]message.Message)
}
