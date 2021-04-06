package main

import "fmt"

func main() {

	var a = 5
	var b = 2

	fmt.Println(a + b) // 덧셈
	fmt.Println(a - b) // 뺄셈
	fmt.Println(a * b) // 곱셈
	fmt.Println(a / b) // 몫
	fmt.Println(a % b) // 나머지

	fmt.Printf("a&b = %v\n", a&b)    // and
	fmt.Printf("result = %v\n", a|b) // or

	var c = 21
	var d = c % 10
	c = c / 10
	var e = c % 10

	fmt.Printf("첫번째 수 : %v 두번째 수 : %v\n", d, e)

	fmt.Println(a << 1) // 한 비트 왼쪽으로 이동
	fmt.Println(a >> 1) // 한 비트 오른쪽으로 이동
}
