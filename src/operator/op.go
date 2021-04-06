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

	fmt.Printf("a&b = %v\n", a&b) // and
	fmt.Printf("a|b = %v\n", a|b) // or
	fmt.Printf("a^b = %v\n", a^b) // xor
	fmt.Println(b << 1)           // 한 비트 왼쪽으로 이동
	fmt.Println(b >> 1)           // 한 비트 오른쪽으로 이동

	fmt.Println(a > b)              // true
	fmt.Println(a <= b)             // false
	fmt.Println(a != b)             // true
	fmt.Println(a > b && a == 5)    // true
	fmt.Println(a < b || a == 5)    // true
	fmt.Println(a < b || !(a == 5)) // false
}
