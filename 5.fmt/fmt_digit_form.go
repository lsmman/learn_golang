package main

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 123456789

	var fa = 324.13455
	var fb = 324.13455
	var fc = 3.14

	// integer 출력 너비 지정
	fmt.Printf("%5d, %5d\n", a, b)    // ❶ 최소 너비보다 짧은 값 너비 지정
	fmt.Printf("%05d, %05d\n", a, b)  // ❷ 최소 너비보다 짧은 값 0 채우기
	fmt.Printf("%-5d, %-05d\n", a, b) // ❸ 최소 너비보다 짧은 값 왼쪽 정렬

	// float 출력 너비 지정
	fmt.Printf("%5d, %5d\n", c, c)    // ➍ 최소 너비보다 긴 값 너비 지정
	fmt.Printf("%05d, %05d\n", c, c)  // ➎ 최소 너비보다 긴 값 0 채우기
	fmt.Printf("%-5d, %-05d\n", c, c) // ➏ 최소 너비보다 긴 값 왼쪽 정렬

	fmt.Printf("%08.2f\n", fa) // ❶ 최소너비:8 소수점이하:2자리 0을 채움.
	fmt.Printf("%08.2g\n", fb) // ❷ 최소너비:8 총숫자: 2자리 0을 채움
	fmt.Printf("%8.5g\n", fb)  // ❸ 최소너비:8 총숫자: 5자리
	fmt.Printf("%f\n", fc)     // ❹ 소수점이하 6자리까지 출력
}
