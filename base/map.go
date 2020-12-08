package main

import "fmt"

type person struct {
	n int
}

func main() {

	var m map[person]person
	p := person{}
	//　不能寻址,不能赋值
	//m[p].n = 1

	fmt.Println(m[p].n)
}
