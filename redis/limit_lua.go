package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

// redis实现滑动窗口
func snowflake() {

	var luatxt = `
local ratelimit_info = redis.pcall('HMGET',KEYS[1],'last_time','current_token')

local last_time = ratelimit_info[1]

local current_token = tonumber(ratelimit_info[2])

local max_token = tonumber(ARGV[1])

local token_rate = tonumber(ARGV[2])

local current_time = tonumber(ARGV[3])

local reverse_time = 1000/token_rate

local result = 0

if current_token == nil then

  current_token = max_token

  last_time = current_time

else

  local past_time = current_time-last_time
  local reverse_token = math.floor(past_time/reverse_time)

  if(reverse_token == 0 and current_token==0 ) then
	return result
  end 
  current_token = current_token+reverse_token

  last_time = reverse_time*reverse_token+last_time

  if current_token>max_token then

    current_token = max_token

  end

end


if(current_token>0) then

  result = 1

  current_token = current_token-1

end 

redis.call('HMSET',KEYS[1],'last_time',last_time,'current_token',current_token)

redis.call('pexpire',KEYS[1],math.ceil(reverse_time*(max_token-current_token)+(current_time-last_time)))

return result`

	var wait sync.WaitGroup
	p := dial("", "")
	for i := 0; i < 8; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			fmt.Println(redis.Int64(p.Get().Do("eval", luatxt, 1, "limt_test", 5, 5, time.Now().UnixNano()/1000000)))
		}()
	}
	wait.Wait()
	fmt.Println(redis.StringMap(p.Get().Do("hgetall", "limt_test")))

}

//incrlua 自增1
func incrlua() {
	var luatxt = `
local limit = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local current = redis.call("GET", KEYS[1])

if current ~= false and tonumber(current) >= limit then
	return 0
end

local current = redis.call("INCRBY", KEYS[1], 1)
if current == 1 then
    redis.call("expire", KEYS[1], window)
    return 1
elseif current > limit then
    return 0
else
    return 1
end`

	var wait sync.WaitGroup
	p := dial("", "")
	for i := 0; i < 8; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			fmt.Println(redis.Int64(p.Get().Do("eval", luatxt, 1, "limt_test", 5, 1)))
		}()
	}
	wait.Wait()
	fmt.Println(redis.String(p.Get().Do("get", "limt_test")))

}
