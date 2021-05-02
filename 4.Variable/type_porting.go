package main

import "fmt"
/*
타입명(변수)로 타입변환 가능

주의사항
1. float를 int로 변환하면 소수점 삭제
2. 더 큰 타입형을 더 작은 타입으로 변환하면 값이 변조
	ex) var x int16 = 3456
		var y int8 = int8(x)
		x와 y의 값은 다르다
*/		
func main() {
	a := 3              // int
	var b float64 = 3.5 // float64

	var c int = int(b)  // ❶ float64에서 int로 변환
	d := float64(a * c) // int에서 float6로 변환

	var e int64 = 7
	f := int64(d) * e // float64에서 int64로 변환

	var g int = int(b * 3) // float64에서 int로 변환
	var h int = int(b) * 3 // float64에서 int로 변환 g와 값이 다릅니다.
	
	fmt.Println(g, h, f)

	// 2. 더 큰 타입형을 더 작은 타입으로 변환하면 값이 변조
	var x int16 = 3456
	var y int8 = int8(x)

	fmt.Println(x, y)
}
