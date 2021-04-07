package main

import "fmt"

func main() {

	var a = 5

	// if a == 3 || a == 4 {
	// 	fmt.Println(("a는 3 또는 4"))
	// } else {
	// 	fmt.Println(("a는 3도 아니고 4도 아니다."))
	// }

	if a == 3 {
		fmt.Println("a는 3")
	} else if a == 4 {
		fmt.Println("a는 4")
	} else {
		fmt.Println(("a는 3도 아니고 4도 아니다."))
	}

	fmt.Println("a의 값은 ", a)
}
