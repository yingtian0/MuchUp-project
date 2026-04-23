package state_kvs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net"
	"sync"

	hc "github.com/Code-Hex/go-generics-cache"
	"github.com/redis/go-redis/v9"
)

type Cache[T any] interface {
	Get(ctx context.Context, key string) (*T, error)
	Set(ctx context.Context, key string, val T) error
	Del(ctx context.Context, key string) error
}

type InMemory[T any] struct {
	client *redis.Client
	local  *hc.Cache[string, T]
	mu     sync.Mutex
	rmu    sync.RWMutex
}

var ErrNotFound = errors.New("Not Found")

var _ Cache[any] = (*InMemory[any])(nil)

func MewInMemory[T any](clinet *redis.Client) *InMemory[T] {
	return &InMemory[T]{
		client: clinet,
		local:  hc.New[string, T](),
		mu:     sync.Mutex{},
		rmu:    sync.RWMutex{},
	}
}

func (m *InMemory[T]) Get(ctx context.Context, key string) (*T, error) {
	var val T
	m.rmu.RLock()
	defer m.rmu.Unlock()

	if v, ok := m.local.Get(key); !ok {
		log.Println("local cache miss")
	} else {
		return &v, nil
	}

	v, err := m.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if err := json.NewDecoder(bytes.NewBufferString(v)).Decode(val); err != nil {
		return nil, err
	}
	m.local.Set(key, val)
	return &val, nil
}

func (m *InMemory[T]) Set(ctx context.Context, key string, val T) error {
	buf := bytes.NewBuffer(nil)
	m.mu.Lock()
	defer m.mu.Unlock()

	if err := json.NewEncoder(buf).Encode(val); err != nil {
		return err
	}
	if _, err := m.client.XAdd(ctx, key, buf.String(), 0).Result(); err != nil {
		if _, ok := err.(net.Error); ok {
			log.Println("Failed to cache %s, continuing without cache", key)

		}
		return err
	}
	return nil

}

func (m *InMemory[T]) Del(ctx context.Context, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.local.Get(key); !ok {
		log.Println("local cache miss")
	} else {
		m.local.Delete(key)
	}

	if _, err := m.client.Del(ctx).Result(); err != nil {
		return err
	}
	return nil

}
