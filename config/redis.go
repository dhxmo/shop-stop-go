package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

const (
	RedisExpiredTimes = 600
)

type Cache interface {
	Get(key string, data interface{}) error
	Set(key string, val []byte) error
	Remove(key string) error
}

type inMemoryCache struct {
	data sync.Map
	mu   sync.Mutex
}

func NewInMemoryCache() Cache {
	return &inMemoryCache{data: sync.Map{}}
}

func (i *inMemoryCache) Get(key string, data interface{}) error {
	val, ok := i.data.Load(key)
	if !ok {
		return nil
	}

	err := json.Unmarshal(val.([]byte), &data)
	if err != nil {
		return err
	}

	return nil
}

func (i *inMemoryCache) Set(key string, val []byte) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.data.Store(key, val)

	return nil
}

func (i *inMemoryCache) Remove(key string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.data.Delete(key)

	return nil
}

type redisCache struct {
	client *redis.Client
}

func NewRedis() *redisCache {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.Database,
	})

	pong, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal(pong, err)
		return nil
	}

	return &redisCache{client: rdb}
}

func (r *redisCache) Get(key string, data interface{}) error {
	val, err := r.client.Get(key).Bytes()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		log.Println("Cache fail to get: ", err)
		return nil
	}
	log.Printf("Get from redis %s - %s", key, val)

	err = json.Unmarshal(val, &data)
	if err != nil {
		return err
	}

	return nil
}

func (r *redisCache) Set(key string, val []byte) error {
	err := r.client.Set(key, val, RedisExpiredTimes*time.Second).Err()
	if err != nil {
		log.Println("Cache fail to set: ", err)
		return err
	}
	log.Printf("Set to redis %s - %s", key, val)

	return nil
}

func (r *redisCache) Remove(key string) error {
	err := r.client.Del(key).Err()
	if err != nil {
		log.Printf("Cache fail to delete key %s: %s", key, err)
		return err
	}
	log.Println("Cache deleted key", key)

	return nil
}
