# Hello World 프로그램은 어떤 과정을 거쳐 실행되는가?

우리는 1번 Hello 코스에서 아래와 같은 프로그램을 작성하였다. 실행하면 상쾌한 Hello world가 출력된다. 

어떤 과정을 거쳐 우리는 Hello World라는 문구를 확인할 수 있게 되는 것일까?
```go
// 패키지 선언
// main은 프로그램의 엔트리 포인트를 포함하는 특별한 패키지
package main

// 표준 입출력을 다루는 내장 패키지인 fmt 패키지 import
import "fmt"

/* main 함수는 프로그램의 첫 진입점
   main() 함수를 선언하고 중괄호 {}로 함수의 시작과 끝을 표시*/
func main(){
    // 표준 출력으로 문자열을 출력하는 Println을 사용하여 hello world 출력
    fmt.Println("Hello World!")
}

// >> Hello world!
```

크게 5가지 과정 (폴더 생성, go 파일 작성, Go 모듈 생성, 빌드, 실행)을 통해 프로그램을 실행된다. 자세히 알아보자. 

```bash
# 폴더 생성
mkdir goprojects
cd goprojects

# go 파일 작성
vi hello.go

# Go 모듈 생성
go mod init goprojects

# 빌드
go build

# 실행
./hello
```
### 1. 폴더(패키지) 생성
- Go 언어에서는 다른 많은 언어들 같이 폴더 단위가 패키지 단위이다. 모든 코드는 폴더(패키지) 단위로 작성된다. 
- 같은 폴더에 위치한 .go 파일은 모두 같은 패키지에 포함되고, 패키지명은 폴더명을 사용해 명명한다.
- 패키지에 대한 추가적인 설명은 16장에서 확인 할 수 있다.



### 2. 파일 .go 작성
- go 확장자를 가진 파일을 생성하여 go 문법에 맞게 원하는 프로그램을 작성한다.


### 3. Go 모듈 생성
- ```go mod init package_name(directory_name)```
- Go 모듈을 생성하면 go.mod 파일이 생성된다. go.mod 파일에는 모듈명과 Go 버전 필요한 패키지 목록 정보가 담겨있다.
- 모듈의 추가적인 설명은 16장에서 확인 할 수 있다.


### 4. 빌드
- go는 정적 컴파일 방법을 채택하고 있기 때문에 c, c++과 같이 실행하기 위해서는 빌드 과정이 필요하다.
- go build 명령은 Go 코드를 기계어로 변환하여 실행 파일을 만든다.
- ```go tool dist list``` 명령을 사용하여 빌드 가능한 운영체제와 아키텍쳐 목록을 사용할 수 있다. 예를 들어 AMD64 계열 칩셋을 사용하는 리눅스 실행 파일을 만들 때는 다음과 같은 옵션을 사용한다. ```(GOOS=linux GOARCH=amd64 go build)```


### 5. 실행
- go build로 생성된 실행 파일을 실행한다.
- 리눅스라면 .o 파일, 윈도우라면 .exe 파일이 생성되어 실행 가능하다.




