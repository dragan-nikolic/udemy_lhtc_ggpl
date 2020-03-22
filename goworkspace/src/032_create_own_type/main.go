package main

import "fmt"

type hotdog int

func main() {
	var a int
	fmt.Println(a)

	var b hotdog

	// a = b  // cannot use b (type hotdog) as type int in assignment
	a = int(b)
	fmt.Println(a)

	fmt.Printf("%T\n", b)
}
