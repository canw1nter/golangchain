package history

import "golangchain/message"

type IHistory interface {
	Add(message []message.Message)
	Get() (*[]message.Message, error)
	Clear()
}
