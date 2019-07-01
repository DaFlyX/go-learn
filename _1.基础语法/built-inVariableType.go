package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
	bool,string
	(u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr
	byte,rune
	浮点数、复数
	float32,float64,complex64,complex128 (3+4i)
*/
//复数运算欧拉公式 1i表示  i
func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1) //(0.000+0.000i)
}

//强制类型转换
//直角定理 a^2+b^2 = c^2
func triangle() {
	var a, b int = 3, 4
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func main() {
	euler()
	triangle()
}
