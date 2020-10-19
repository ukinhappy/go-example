package randx

import (
	"math/rand"
	"strconv"
	"time"
)

//RandomString  前缀为时间戳的随机数目
func TimeWithRandomString(lenth int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := time.Now().Unix() * 100
	n := 10
	for i := 1; i < lenth; i++ {
		n = n * 10
	}
	data = data + int64(r.Intn(n))
	return strconv.Itoa(int(data))
}

func RandomInt(lenth int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := 10
	for i := 1; i < lenth; i++ {
		n = n * 10
	}
	return r.Intn(n)
}
