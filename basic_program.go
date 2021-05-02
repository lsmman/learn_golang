/* 실행방법
1. go mod init package_name (=fold_name)
2. go build 
3. 실행 파일 ex) hello.exe
*/

package main

import "fmt"

func main() {
	var msg string = "Lim"

	fmt.Println("Hello", msg)
}