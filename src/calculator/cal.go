package main

import "fmt"

func main() {
	//번수 선언 따로, 대입 따로
	var a int
	var b int
	a = 3
	b = 4

	//타입 없이 대입
	var d = 6
	var e = 3.14

	fmt.Println(a + b)

	var word string = "abc"
	fmt.Println(word)

	fmt.Println(d)
	fmt.Println(e)
}
