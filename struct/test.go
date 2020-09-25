package main

import (
	"fmt"
)

func main() {
	v1 := struct {
		v int
	}{v: 1}
	v2 := struct {
		v int
	}{v: 2}
	//可以判等
	if v1 == v2 {
		fmt.Println("v1=v2")
	}



	m1 := struct {
		m map[int]int
	}{m: map[int]int{1: 1}}
	m2 := struct {
		m map[int]int
	}{m: map[int]int{1: 1}}

	// 不能判等　结构体比较，成员必须可比较
	if m1 == m2{}
	fmt.Println("m1 = m2")
}
