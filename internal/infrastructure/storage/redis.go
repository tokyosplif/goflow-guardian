package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tokyosplif/goflow-guardian/internal/config"
	"github.com/tokyosplif/goflow-guardian/internal/domain"
)

const (
	redisKeyPrefix = "limiter:"
)

type RedisLimiter struct {
	client *redis.Client
}

func NewRedisLimiter(cfg config.Redis) *RedisLimiter {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return &RedisLimiter{client: rdb}
}

func (r *RedisLimiter) IsAllowed(ctx context.Context, key string, limit domain.Limit) (bool, error) {
	now := time.Now().UnixNano()
	boundary := now - limit.Window.Nanoseconds()
	redisKey := fmt.Sprintf("%s%s", redisKeyPrefix, key)

	pipe := r.client.TxPipeline()
	pipe.ZRemRangeByScore(ctx, redisKey, "0", fmt.Sprintf("%d", boundary))
	count := pipe.ZCard(ctx, redisKey)
	pipe.ZAdd(ctx, redisKey, redis.Z{Score: float64(now), Member: now})
	pipe.Expire(ctx, redisKey, limit.Window)

	if _, err := pipe.Exec(ctx); err != nil {
		return false, err
	}

	return count.Val() < int64(limit.Requests), nil
}

func (r *RedisLimiter) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}
