package main

import "fmt"

/*
	数组
	值传递 调用会拷贝数组
*/

func main() {
	//数组定义
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(arr1, arr2, arr3) //[0 0 0 0 0] [1 3 5] [2 4 6 8 10]
	//二维数组定义
	var grid [4][5]bool
	fmt.Println(grid)
	/*
		[[false false false false false]
		[false false false false false]
		[false false false false false]
		[false false false false false]]
	*/

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
		/*
			2
			4
			6
			8
			10
		*/
	}
	//i 索引 v 值
	for i, v := range arr3 {
		fmt.Println(i, v)
		/*
			0 2
			1 4
			2 6
			3 8
			4 10

		*/
	}
	sum := 0
	numbers := [...]int{10, 9, 8, 7, 6, 5}
	for i := range numbers {
		fmt.Println("i:", i)//索引
		/*
			i: 0
			i: 1
			i: 2
			i: 3
			i: 4
			i: 5
		*/
		sum += numbers[i]
	}
	fmt.Println(sum) //45
}
