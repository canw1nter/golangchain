package history

import (
	"context"
	"time"

	"golangchain/common"
	"golangchain/message"

	"github.com/redis/go-redis/v9"
)

type RedisHistoryOption struct {
	Url    string
	Prefix string
	TTL    time.Duration
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

func (rh *RedisHistory) Add(messages []message.Message) {
	for _, m := range messages {
		rh.client.Set(context.Background(), rh.Prefix+rh.SessionId, m, rh.TTL)
	}
}

func (rh *RedisHistory) Clear() {
	rh.client.Del(context.Background(), rh.Prefix+rh.SessionId)
}

func NewRedisHistory(SessionId string, opts ...common.Options) (*RedisHistory, error) {
	history := &RedisHistory{
		History:   &History{Messages: make([]message.Message, 0)},
		SessionId: SessionId,
		RedisHistoryOption: &RedisHistoryOption{
			Url:    "redis://localhost:6379/0",
			Prefix: "GLC_HISTORY:",
			TTL:    0,
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

func WithTTL(TTL time.Duration) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*RedisHistoryOption); ok {
			options.TTL = TTL
		}
	}
}
