package lockx

import (
	"go-example/utils/randx"
	redis "gopkg.in/redis.v5"
	"sync"
	"time"
)

//ClusterOption
type ClusterOption func(*ClusterOptions)

//ClusterOptions
type ClusterOptions struct {
	// A seed list of host:port addresses of cluster nodes.
	Addrs []string

	// The maximum number of retries before giving up. Command is retried
	// on network errors and MOVED/ASK redirects.
	// Default is 16.
	MaxRedirects int
	// Following options are copied from Options struct.

	Password string

	ReadTimeout time.Duration
	IdleTimeout time.Duration

	// PoolSize applies per cluster node and not for the whole cluster.
	PoolSize int
}

var DefaultClusterOptions ClusterOptions

//WithAddrs
func WithAddrs(addrs []string) ClusterOption {
	return func(ops *ClusterOptions) {
		ops.Addrs = addrs
	}
}

//WithMaxRedirects
func WithMaxRedirects(maxRedirects int) ClusterOption {
	return func(ops *ClusterOptions) {
		ops.MaxRedirects = maxRedirects
	}
}

//WithPassword
func WithPassword(passWord string) ClusterOption {
	return func(ops *ClusterOptions) {
		ops.Password = passWord
	}
}

var defaultCluster *redis.ClusterClient
var once sync.Once

// Init 初始化redis
func InitCluster(opts ...ClusterOption) {
	var option = DefaultClusterOptions
	for _, opt := range opts {
		opt(&option)
	}

	once.Do(func() {
		defaultCluster = newCluster(option)
	})
}

func newCluster(redisConf ClusterOptions) (cluster *redis.ClusterClient) {

	opts := redis.ClusterOptions{
		Addrs:        redisConf.Addrs,
		MaxRedirects: redisConf.MaxRedirects,
		Password:     redisConf.Password,
		PoolSize:     redisConf.PoolSize,
		ReadTimeout:  500 * time.Millisecond,
		IdleTimeout:  time.Second * 240,
	}
	cluster = redis.NewClusterClient(&opts)
	return
}

//Mutex mutex
type (
	Mutex struct {
		ResourceName string
		Token        string
	}
)

//NewMutex new mutex
func NewMutex(resourceName string) *Mutex {
	return &Mutex{
		Token:        randx.TimeWithRandomString(4),
		ResourceName: "{im:lockx}" + resourceName,
	}
}

//Lock mutex lockx
func (m *Mutex) Lock(lockTime int64) bool {
	result := defaultCluster.SetNX(m.ResourceName, m.Token, time.Duration(lockTime)*time.Second)
	if err := result.Err(); err != nil {
		return false
	}
	return result.Val()
}

//UnLock mutex unlockw
func (m *Mutex) UnLock() bool {
	result := defaultCluster.Eval(m.luaScripts(), []string{m.ResourceName}, m.Token)
	if err := result.Err(); err != nil {
		return false
	}
	if v, ok := result.Val().(int64); ok {
		if v == 1 {
			return true
		} else {
			return false
		}
	}
	return false
}

func (m *Mutex) luaScripts() string {
	lua := `
if redis.call("get",KEYS[1]) == ARGV[1] 
then
    return redis.call("del",KEYS[1])
else
    return 0
end`
	return lua
}
