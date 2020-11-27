package reflectx

import (
	"go-example/utils/timex"
	"testing"
	"time"
)

func TestObjectToMap(t *testing.T) {

	type User struct {
		Name     string `json:"name"`
		Age      int
		Age1     int32   `json:"age1"`
		Age2     float64 `json:"age2"`
		Age3     uint32  `json:"age3"`
		Age4     uint64
		Age5     float32
		Birthday time.Time
	}
	happy := &User{
		"happy", 1, 1, 1.0, 1, 1, 1, time.Now()}
	t.Log(ObjectToMap(happy, "json"))
}

func TestMapToObject(t *testing.T) {

	type User struct {
		Name string `json:"name"`
		Age  int
		Age1 int32   `json:"age1"`
		Age2 float64 `json:"age2"`
		Age3 uint32  `json:"age3"`
		Time time.Time
	}

	value := ObjectToMap(&User{
		"happy", 1, 33333, 1.0, 1, timex.Now().T}, "json")
	var happy User
	MapToObject(&happy, value, "json")
	t.Log(happy)
}
