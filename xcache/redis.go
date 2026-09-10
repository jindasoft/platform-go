package xcache

import (
	"context"
	"fmt"
	"time"

	"github.com/jindasoft/template-platform-go/xlogger"
	"github.com/redis/go-redis/v9"
)

type RedisService interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	SetNX(ctx context.Context, key string, value any, expiration time.Duration) (bool, error)
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Incr(ctx context.Context, key string) (int64, error)
	Decr(ctx context.Context, key string) (int64, error)
	HSet(ctx context.Context, key string, field string, value any) error
	HGet(ctx context.Context, key string, field string) (string, error)
	Client() *redis.Client
	Health(ctx context.Context) HealthStats
}

type RedisConfig struct {
	Host         string
	Port         int
	Password     string `json:"-"`
	InstanceName string
}

type service struct {
	client *redis.Client
}

type HealthStats struct {
	Code      string
	Message   string
	Error     error
	PoolStats *redis.PoolStats
}

const (
	keyFormat = "%s:%s"
)

var (
	instanceName string
)

func NewRedis(ctx context.Context, cfg *RedisConfig) (*service, error) {
	instanceName = cfg.InstanceName

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		xlogger.SysErrorf("Failed to ping Redis: %v", err)
		return nil, err
	}

	xlogger.SysInfof("Redis initialized successfully")
	return &service{client}, nil
}

func (s *service) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	k := fmt.Sprintf(keyFormat, instanceName, key)
	return s.client.Set(ctx, k, value, expiration).Err()
}

func (s *service) Get(ctx context.Context, key string) (string, error) {
	k := fmt.Sprintf(keyFormat, instanceName, key)
	return s.client.Get(ctx, k).Result()
}

func (s *service) SetNX(ctx context.Context, key string, value any, expiration time.Duration) (bool, error) {
	k := fmt.Sprintf(keyFormat, instanceName, key)
	args := redis.SetArgs{
		Mode: "NX",
		TTL:  expiration,
	}
	// res, err := s.client.SetNX(ctx,k,value, expiration).Result()
	res, err := s.client.SetArgs(ctx, k, value, args).Result()
	return res == "OK", err
}

func (s *service) Delete(ctx context.Context, key string) error {
	k := fmt.Sprintf(keyFormat, instanceName, key)
	return s.client.Del(ctx, k).Err()
}

func (s *service) Client() *redis.Client {
	return s.client
}

func (s *service) Incr(ctx context.Context, key string) (int64, error) {
	k := fmt.Sprintf(keyFormat, instanceName, key)
	return s.client.Incr(ctx, k).Result()
}

func (s *service) Decr(ctx context.Context, key string) (int64, error) {
	k := fmt.Sprintf(keyFormat, instanceName, key)
	return s.client.Decr(ctx, k).Result()
}

func (s *service) HSet(ctx context.Context, key string, field string, value any) error {
	k := fmt.Sprintf(keyFormat, instanceName, key)
	return s.client.HSet(ctx, k, field, value).Err()
}

func (s *service) HGet(ctx context.Context, key string, field string) (string, error) {
	k := fmt.Sprintf(keyFormat, instanceName, key)
	return s.client.HGet(ctx, k, field).Result()
}

func (s *service) Health(ctx context.Context) HealthStats {
	stats := HealthStats{}
	if err := s.client.Ping(ctx).Err(); err != nil {
		stats.Code = "down"
		stats.Error = fmt.Errorf("redis down: %v", err)
		xlogger.SysErrorf("Redis health check failed: %v", err)
		return stats
	}
	stats.Code = "up"
	stats.Message = "Redis is healthy"
	stats.PoolStats = s.client.PoolStats()
	if stats.PoolStats.TotalConns > 100 {
		stats.Message = "Redis is experiencing high connection usage."
	}
	return stats
}
