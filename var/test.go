package main

type Int1 int

func main() {

	var i int = 0
	// 不能赋值
	var i1 Int1 = i

	var x = nil
	var y interface{} = nil
	var z string = nil
	var e error = nil
}
