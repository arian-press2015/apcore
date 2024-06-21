package keystore

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type KeyStore struct {
	client *redis.Client
}

func NewKeyStore(addr string) *KeyStore {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &KeyStore{client: client}
}

func (r *KeyStore) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *KeyStore) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *KeyStore) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.client.SetNX(ctx, key, value, expiration).Result()
}

func (r *KeyStore) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}
