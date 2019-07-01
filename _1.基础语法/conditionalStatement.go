package main

import (
	"fmt"
	"io/ioutil"
)

/*
	条件语句 if switch
*/
//使用if条件语句
func If() {
	const filename = "_1.基础语法/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("读取内容:", string(contents))
	}
}

//使用switch语句
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}
func main() {
	If()
	/*
		读取内容: 1
		2
		3
		4
	*/
	fmt.Println(
		eval(1, 2, "+"), //3
		eval(1, 2, "-"), //-1
		eval(1, 2, "*"), //2
		eval(1, 2, "/"), //0
		//eval(1, 2, "//"),
	)
	/*
		没有匹配到就会报错一旦panic,程序停止运行
		panic: unsupported operator://

		goroutine 1 [running]:
		main.eval(...)
			D:/GoWork/src/go-learn/_1.基础语法/conditionalStatement.go:34
		main.main()
			D:/GoWork/src/go-learn/_1.基础语法/conditionalStatement.go:51 +0x178

	*/

}
