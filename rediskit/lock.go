package rediskit

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/go-redis/redis"
)

// RedisLock Redis distributed lock
type RedisLock struct {
	c     *redis.Client
	key   string
	value string
	ch    chan struct{}
}

// NewRedisLock new a redis lock
// key - lock name
func NewRedisLock(c *redis.Client, key string) *RedisLock {
	return &RedisLock{
		c:   c,
		key: key,
		ch:  make(chan struct{}),
	}
}

// Lock try lock
// notice: renew every 15 sec
func (l *RedisLock) Lock() error {
	var guid [16]byte
	_, err := io.ReadFull(rand.Reader, guid[:16])
	if err != nil {
		return err
	}
	guidStr := fmt.Sprintf("%x", guid)

	l.value = guidStr
	ok, err := l.c.SetNX(l.key, l.value, 30*time.Second).Result()
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("can not lock")
	}

	go func() {
		ticker := time.NewTicker(15 * time.Second)
		for {
			select {
			case <-ticker.C:
				_ = l.c.Expire(l.key, 30*time.Second).Err()

			case <-l.ch:
				return
			}
		}
	}()

	return nil
}

// Unlock unlock check value
func (l *RedisLock) Unlock() error {
	value, err := l.c.Get(l.key).Result()
	if err != nil {
		return err
	}

	if value != l.value {
		return errors.New("can not match lock value")
	}

	_, err = l.c.Del(l.key).Result()
	if err != nil {
		return err
	}

	l.ch <- struct{}{}
	return nil
}
