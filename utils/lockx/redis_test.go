package lockx

import (
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestRedis(t *testing.T) {
	InitSyncLock(redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379", nil)
	}, 1))
	NewSyncLocker("lock", redsync.SetExpiry(1), redsync.SetRetryDelay(1))
}
