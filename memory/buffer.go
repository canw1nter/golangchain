package memory

import "golangchain/message"

type BufferMemory struct {
	*Memory
}

func (bm *BufferMemory) GetMemory() []message.Message {
	//TODO implement me
	panic("implement me")
}

func (bm *BufferMemory) ClearMemory() {
	//TODO implement me
	panic("implement me")
}

func (bm *BufferMemory) SaveToMemory(messages []message.Message) {
	//TODO implement me
	panic("implement me")
}
