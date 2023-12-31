package history

import (
	"context"
	"encoding/json"
	"time"

	"golangchain/common"
	"golangchain/message"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type RedisHistoryOption struct {
	Url    string
	Prefix string
	TTL    time.Duration
}

type RedisHistory struct {
	client    *redis.Client
	SessionId string
	*RedisHistoryOption
}

func (rh *RedisHistory) SetOptions(opts ...common.Options) {
	for _, opt := range opts {
		opt(rh.RedisHistoryOption)
	}
}

func (rh *RedisHistory) Get() (*[]message.Message, error) {
	messages := make([]message.Message, 0)

	data, err := rh.client.LRange(context.Background(), rh.Prefix+rh.SessionId, 0, -1).Result()
	if err != nil {
		return nil, errors.Wrap(err, "get history failed")
	}

	if len(data) != 0 {
		for _, v := range data {
			var m message.Message
			err = json.Unmarshal([]byte(v), &m)
			if err != nil {
				return nil, errors.Wrap(err, "get history failed")
			}

			messages = append(messages, m)
		}
	}

	return &messages, nil
}

func (rh *RedisHistory) Add(messages []message.Message) {
	rh.client.LPush(context.Background(), rh.Prefix+rh.SessionId, messages)
	rh.client.Expire(context.Background(), rh.Prefix+rh.SessionId, rh.TTL)
}

func (rh *RedisHistory) Clear() {
	rh.client.Del(context.Background(), rh.Prefix+rh.SessionId)
}

func NewRedisHistory(SessionId string, opts ...common.Options) (*RedisHistory, error) {
	history := &RedisHistory{
		SessionId: SessionId,
		RedisHistoryOption: &RedisHistoryOption{
			Url:    "redis://localhost:6379/0",
			Prefix: "GLC_HISTORY:",
			TTL:    0,
		},
	}

	history.SetOptions(opts...)

	redisOpt, err := redis.ParseURL(history.Url)
	if err != nil {
		return nil, errors.Wrap(err, "create new redis history failed")
	}
	history.client = redis.NewClient(redisOpt)

	return history, nil
}

func WithUrl(url string) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*RedisHistoryOption); ok {
			options.Url = url
		}
	}
}

func WithPrefix(prefix string) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*RedisHistoryOption); ok {
			options.Prefix = prefix
		}
	}
}

func WithTTL(ttl time.Duration) common.Options {
	return func(obj interface{}) {
		if options, ok := obj.(*RedisHistoryOption); ok {
			options.TTL = ttl
		}
	}
}
