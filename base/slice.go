package main

import "fmt"

func main()  {

	//  [0,0,0,0,0,1,2,3]
	s1:= make([]int,5)
	s1= append(s1,1,2,3)
	fmt.Println(s1)

	// [1,2,3,4,5]
	s2:=make([]int,0)
	s2 = append(s2,1,2,3,4,5)
	fmt.Println(s2)

	//// 不能编译通过 s3是个 *[]int　不能向里append
	//s3:= new([]int)
	//s3= append(s3,1)

	hello(s1...)
	fmt.Println(s1)
	s1 =[]int{1,2,3,4,5}
	var r [5]int
	for i, v := range s1 {
		if i == 0 {
			s1[1] = 12
			s1[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", s1)
}

func hello(s ...int)  {
	s[0]=10
}