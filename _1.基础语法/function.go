package main

import (
	"fmt"
	"strconv"
)

/*
	函数定义
*/
//定义函数
func grade(score int) string {
	result := ""
	switch {
	case score >= 90 && score <= 100:
		result = "A"
	case score >= 80 && score < 90:
		result = "B"
	case score >= 70 && score < 80:
		result = "C"
	case score >= 60 && score < 70:
		result = "D"
	default:
		result = "F"
	}
	return result
}

//返回多个值 推荐 第一种
func div(a, b int) (int, int) {
	return a / b, a % b
}
func div1(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}
//函数作为参数
func apply(grade func(int) string, a int) string {
	return grade(a)
}
//可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers{
		fmt.Println(i)
		s += numbers[i]
	}
	return s
}
func main() {
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(79),
		grade(89),
		grade(99),
	)
	/*
		F F D C B A
	*/
	q, r := div(13, 3)
	fmt.Println(q, r) //4 1
	a, b := div1(-13, 3)
	fmt.Println(a, b) //-4 -1
	c, d := div1(-1, 3)
	fmt.Println(c, d) //0 -1

	fmt.Println(
		apply(grade, 10),
	) //F

	fmt.Println(apply(
		func(i int) string {
			return fmt.Sprintf("%q\n",strconv.Itoa(i))
		},
		10),
	)//"10"
	fmt.Println(sum(1,2,3,4,5)) //15
}
