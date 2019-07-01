package main

import (
	"fmt"
	"math"
)

/*
	常量与枚举
*/
//定义常量
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c) //abc.txt 5

}
//定义枚举
func enums() {
	const (
		cpp = iota
		python
		java
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, python, java, golang, javascript)//0 1 2 3 4
	fmt.Println(b, kb, mb, gb, tb, pb)//1 1024 1048576 1073741824 1099511627776 1125899906842624
}

func main(){
	consts()
	enums()
}