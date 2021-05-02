package main

import "fmt"

var g int = 10 // 패키지 전역 변수 선언 ❶

func main() {
	var m int = 20 // 로컬 변수 선언 ❷

	{
		var s int = 50 // 로컬 변수 선언 ❸
		fmt.Println(m, s, g)
	} // s 로컬변수는 사라짐 - ❹

	m = s + 20 // Error - ➎
}