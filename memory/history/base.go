package history

import "golangchain/message"

type History struct {
	Messages []message.Message
}

type IHistory interface {
	Add(message []message.Message)
	Clear()
}
