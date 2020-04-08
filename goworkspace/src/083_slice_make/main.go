package main

import "fmt"

func main() {
	x := []int{4, 5, 42, 11}
	fmt.Println(x)

	y := make([]int, 10, 12)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
	//fmt.Println(y[10])  // generates exception

	y = append(y, 99)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	y = append(y, 30)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	// at this point capacity will double (because we went over current capacity)
	y = append(y, 67)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
}
