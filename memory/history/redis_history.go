package history

import (
	"golangchain/common"
	"golangchain/message"

	"github.com/redis/go-redis/v9"
)

type RedisHistoryOption struct {
	Url    string
	Prefix string
	TTL    int
}

type RedisHistory struct {
	*History
	client    *redis.Client
	SessionId string
	*RedisHistoryOption
}

func (rh *RedisHistory) SetOptions(opts ...common.Options) {
	for _, opt := range opts {
		opt(rh.RedisHistoryOption)
	}
}

func (rh *RedisHistory) Add(message []message.Message) {
	//TODO implement me
	panic("implement me")
}

func (rh *RedisHistory) Clear() {
	//TODO implement me
	panic("implement me")
}

func NewRedisHistory(SessionId string, opts ...common.Options) (*RedisHistory, error) {
	history := &RedisHistory{
		History:   &History{Messages: make([]message.Message, 0)},
		SessionId: SessionId,
		RedisHistoryOption: &RedisHistoryOption{
			Url:    "redis://localhost:6379/0",
			Prefix: "GLC_HISTORY:",
			TTL:    3600,
		},
	}

	redisOpt, err := redis.ParseURL(history.Url)
	if err != nil {
		return nil, err
	}
	history.client = redis.NewClient(redisOpt)

	history.SetOptions(opts...)

	return history, nil
}

func WithUrl(Url string) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*RedisHistoryOption); ok {
			options.Url = Url
		}
	}
}

func WithPrefix(Prefix string) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*RedisHistoryOption); ok {
			options.Prefix = Prefix
		}
	}
}

func WithTTL(TTL int) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*RedisHistoryOption); ok {
			options.TTL = TTL
		}
	}
}
