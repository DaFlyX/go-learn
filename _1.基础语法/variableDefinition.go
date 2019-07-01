package main

/*
	定义变量
*/
import "fmt"

/* 一行行定义
var aa = 3
var ss = "kkk"
var bb = true
*/
/*多行定义
var(
	aa = 3
	ss = "kkk"
	bb = true
)*/
//定义变量时  零值
func varZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s) //0 ""
}

//定义变量初始值
func varInitValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s) //3 4 abc
}

//定义变量 编译器推断类型
func varTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s) // 3 4 true def
}

//定义变量 简短定义方式
func varShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 6
	fmt.Println(a, b, c, s) //3 6 true def
}
func main() {
	//调用
	varZeroValue()
	varInitValue()
	varTypeDeduction()
	varShorter()
}
