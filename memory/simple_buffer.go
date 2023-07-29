package memory

import "golangchain/message"

type BufferMemory struct {
	*Memory
}

func (bm *BufferMemory) GetMemory() (*[]message.Message, error) {
	return bm.Memory.GetMemory()
}

func (bm *BufferMemory) ClearMemory() {
	bm.Memory.ClearMemory()
}

func (bm *BufferMemory) SaveToMemory(messages []message.Message) {
	bm.Memory.SaveToMemory(messages)
}
