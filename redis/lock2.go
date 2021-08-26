package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

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
		Token:        "",
		ResourceName: "{lock}" + resourceName,
	}
}

//Lock mutex lock
func (m *Mutex) Lock(lockTime int64) bool {
	p := dial("", "")
	result, err := redis.Bool(p.Get().Do("SetNX", m.ResourceName, m.Token, time.Duration(lockTime)*time.Second))
	if err != nil {
		return false
	}
	return result
}

//UnLock mutex unlock
func (m *Mutex) UnLock() bool {
	p := dial("", "")
	result, err := redis.Int(p.Get().Do("Eval", m.luaScripts(), []string{m.ResourceName}, m.Token))
	if err != nil {
		return false
	}
	if result == 1 {
		return true
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
