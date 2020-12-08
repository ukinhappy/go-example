package main

import "fmt"

const (
	x = iota
	_
	a
	y
	z = "zzz"
	k
	p = iota
)

func main() {

	fmt.Println(x, a, y, z, k, p)
}
