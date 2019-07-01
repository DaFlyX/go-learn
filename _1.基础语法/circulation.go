package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	循环 for

*/
//int 转 二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//只有部分条件
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}
func main() {
	fmt.Println(
		convertToBin(5),       //0101
		convertToBin(13),      //1101
		convertToBin(2321412), //1000110110110000000100
		convertToBin(0),       //
	)
	printFile("_1.基础语法/abc.txt")
	/*
		1
		2
		3
		4
	*/
	forever()
	/*
		abc
		abc
		abc
		abc
		abc
		abc
		abc
		abc
		abc
		abc
		abc
		...
	*/
}
