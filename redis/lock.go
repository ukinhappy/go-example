package main

import (
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

var (
	rs *redsync.Redsync
)

type Pool struct {
	*redis.Pool
}

func (p *Pool) Get() redis.Conn {
	return p.Pool.Get()
}

//InitSyncLock InitSyncLock
func InitSyncLock(p *redis.Pool) {
	var p1 redsync.Pool = &Pool{
		Pool: p,
	}
	rs = redsync.New([]redsync.Pool{p1})
	return
}

//NewSyncLocker NewSyncLocker
func NewSyncLocker(name string, options ...redsync.Option) RedSyncLocker {
	return RedSyncLocker{mu: rs.NewMutex(name, options...)}
}

//RedSyncLocker RedSyncLocker
type RedSyncLocker struct {
	mu *redsync.Mutex
}

//Lock Lock
func (locker *RedSyncLocker) Lock() error {
	return locker.mu.Lock()
}

//Unlock Unlock
func (locker *RedSyncLocker) Unlock() error {
	locker.mu.Unlock()
	return nil
}
