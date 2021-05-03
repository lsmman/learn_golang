package main

import "fmt"

func main() {
	var a int
	var b int
	var f float32 = 123451242.12
	var str string = "Hello World"

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println("n, a, b : ", n, a, b)
	}

	fmt.Println("f :", f)
	fmt.Printf("a, b, f : %d, %d, %f\n", a, b, f)
	fmt.Printf("%s\n", str)
	fmt.Printf("%q\n", str)
}
