package main

import "fmt"

func main()  {

	// 1. for 循环局部变量引用问题
	slice := []int{1,2,3,4}
	m := make(map[int]*int)
	for k,v:=range slice{
		// 不同的ｋ保存的都是同一个ｖ的地址
		m[k]=&v
	}
	for k,v:=range m{
		fmt.Println(k,v,*v)
	}

}
