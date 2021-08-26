package parallel

import (
	"github.com/mohae/deepcopy"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"context"
	"reflect"
	"sync"
)

func copyArgs(args interface{}) interface{} {
	return deepcopy.Copy(args)
}


const concurrencyLimit = 4

//ForEach 并行的迭代数组.
//
// arr 数组.
//
// f 处理函数.
//
// limit 最大并行goroutine数.
func ForEach(arr interface{},
	f func(i int, v interface{}) (interface{}, error), limit int) func() ([]interface{}, error) {
	return func() ([]interface{}, error) {
		varr := reflect.ValueOf(arr)
		ret := make([]interface{}, varr.Len())

		if limit <= 0 {
			limit = concurrencyLimit
		}
		if limit > varr.Len() {
			limit = varr.Len()
		}

		sema := semaphore.NewWeighted(int64(limit))
		g := errgroup.Group{}

		for i := 0; i < varr.Len(); i++ {
			sema.Acquire(context.TODO(), 1)
			ii := i
			v := varr.Index(i).Interface()
			g.Go(func() error {
				defer sema.Release(1)
				d, err := f(ii, v)
				ret[ii] = d
				return err
			})
		}

		err := g.Wait()
		return ret, err
	}
}

//ForEachMap 并行的迭代map
//
// obj map.
//
// f 处理函数.
//
// limit 最大并行goroutine数.
func ForEachMap(obj interface{},
	f func(k, v interface{}) (interface{}, error), limit int) func() (map[interface{}]interface{}, error) {
	return func() (map[interface{}]interface{}, error) {
		vobj := reflect.ValueOf(obj)
		ret := make(map[interface{}]interface{}, vobj.Len())


		var mu sync.Mutex
		if limit <= 0 {
			limit = concurrencyLimit
		}
		if limit > vobj.Len() {
			limit = vobj.Len()
		}

		sema := semaphore.NewWeighted(int64(limit))
		g := errgroup.Group{}

		keys := vobj.MapKeys()
		for _, k := range keys {
			v := vobj.MapIndex(k)
			sema.Acquire(context.TODO(), 1)

			ik := k.Interface()
			iv := v.Interface()
			g.Go(func() error {
				defer sema.Release(1)

				d, err := f(ik, iv)
				//加锁防止多个goroutine对map操作导致的panic
				mu.Lock()
				ret[ik] = d
				mu.Unlock()

				return err
			})
		}

		err := g.Wait()
		return ret, err
	}
}



//Parallel 并行的运行函数并返回函数结果
//
// functions 处理函数数组.
func Parallel(functions []func(...interface{}) (interface{}, error)) func(args []interface{}) ([]interface{}, error) {
	return func(args []interface{}) ([]interface{}, error) {
		if args != nil && len(functions) != len(args) {
			panic("args len do not match functions len")
		}
		ret := make([]interface{}, len(functions))
		sema := semaphore.NewWeighted(int64(len(functions)))
		g := errgroup.Group{}

		for i := 0; i < len(functions); i++ {
			ii := i
			sema.Acquire(context.TODO(), 1)

			var cpy interface{}
			if args != nil {
				cpy = copyArgs(args[i])
			}

			g.Go(func() error {
				defer sema.Release(1)

				d, err := functions[ii](cpy)
				ret[ii] = d

				return err
			})
		}

		err := g.Wait()
		return ret, err
	}
}
