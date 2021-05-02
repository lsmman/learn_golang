package main

import "fmt"

/* 변수 선언 방법 4가지
기본 : 변수 변수이름 타입순으로 선언
초깃값 생략 : 이 때 a는 기본값 0으로 assign
타입 생략 : 타입을 선언하지 않고 assign하면 
선언대입 : var, 타입 선언을 :=을 이용해서 생략
*/

func main(){
	// 변수 선언 방법 4가지
	var a int = 28 // 기본
	var b int  // 초깃값 생략
	var c = 100 // 타입생략
	d := 5 // 선언대입

	var sum_val int = a * c

	fmt.Println("a, b, c, d : ", a, b, c, d)
	fmt.Println("Sum_val = ", sum_val)
}