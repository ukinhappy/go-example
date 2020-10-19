package lockx

import (
	"github.com/go-redsync/redsync"
)

var (
	rs *redsync.Redsync
)

//InitSyncLock InitSyncLock
func InitSyncLock(p redsync.Pool) {
	once.Do(func() {
		rs = redsync.New([]redsync.Pool{p})
	})
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
