package main

import (
	"bufio" // reat string
	"fmt"
	"os"      // 입력
	"strconv" // 문자열 -> 숫자
	"strings" // 찌꺼기 제거
)

func main() {

	fmt.Println("숫자를 입력하세요.")

	reader := bufio.NewReader(os.Stdin)
	// 한 줄을 읽는다. _는 에러를 처리하지 않는 이름없는 변수
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line) // 개행문자와 같은 찌꺼기를 날림

	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다\n", n1, n2)

	fmt.Println("연산자를 입력하세요")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	// if line == "+" {
	// 	fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	// } else if line == "-" {
	// 	fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	// } else if line == "*" {
	// 	fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	// } else if line == "/" {
	// 	fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	// } else {
	// 	fmt.Println("잘못 입력하셨습니다.")
	// }

	switch line {
	case "+":
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	case "/":
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	default:
		fmt.Println("잘못 입력하셨습니다.")
	}
}
