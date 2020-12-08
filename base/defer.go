package main

import "fmt"

func main()  {
	//a := 1
	//b := 2
	//defer calc("1", a, calc("10", a, b))
	//a = 0
	//defer calc("2", a, calc("20", a, b))
	//b = 1

	fmt.Println(f(3))
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}

