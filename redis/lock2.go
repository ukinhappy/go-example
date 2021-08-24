package main

import (
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
	p := dial()
	result, err := p.Get().Do("SetNX", m.ResourceName, m.Token, time.Duration(lockTime)*time.Second))
	if err != nil {
		return false
	}
	return result.Val()
}

//UnLock mutex unlock
func (m *Mutex) UnLock() bool {
	p := dial()
	result, err := p.Get().Do("Eval", m.luaScripts(), []string{m.ResourceName}, m.Token)
	if err != nil {
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
