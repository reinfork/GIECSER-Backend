// Package store holds the Redis cache: PIN metadata + rate-limit counters.
package store

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

// ErrCacheMiss is returned when a key is absent — callers fall back to Postgres.
var ErrCacheMiss = errors.New("cache miss")

type Redis struct {
	c *redis.Client
}

func NewRedis(rawurl string) *Redis {
	// Accepts full URLs (redis://[:password@]host:6379/0) or a bare host:port.
	// ponytail: no sentinel/cluster support — single instance is the whole design.
	opt, err := redis.ParseURL(rawurl)
	if err != nil || rawurl == "" {
		opt = &redis.Options{Addr: "localhost:6379"}
		if err == nil && rawurl != "" {
			opt.Addr = rawurl
		}
	}
	return &Redis{c: redis.NewClient(opt)}
}

func (r *Redis) Ping(ctx context.Context) error {
	return r.c.Ping(ctx).Err()
}

func (r *Redis) SetPIN(ctx context.Context, code, chapterID string, ttl time.Duration) error {
	return r.c.Set(ctx, "pin:"+code, chapterID, ttl).Err()
}

func (r *Redis) GetPIN(ctx context.Context, code string) (string, error) {
	chapterID, err := r.c.Get(ctx, "pin:"+code).Result()
	if errors.Is(err, redis.Nil) {
		return "", ErrCacheMiss
	}
	return chapterID, err
}

// Allow implements a fixed window: limit hits per window per key.
// ponytail: fixed window, not token bucket — upgrade when abuse patterns demand precision.
func (r *Redis) Allow(ctx context.Context, key string, limit int, window time.Duration) (bool, error) {
	n, err := r.c.Incr(ctx, "rl:"+key).Result()
	if err != nil {
		return false, err
	}
	if n == 1 {
		if err := r.c.Expire(ctx, "rl:"+key, window).Err(); err != nil {
			return false, err
		}
	}
	return n <= int64(limit), nil
}
