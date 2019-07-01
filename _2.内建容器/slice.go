package main

import "fmt"

/*
	切片
*/

func updateSlice(s []int) {
	s[0] = 100
}
func printSlice(s1 []int) {
	fmt.Printf("s1=%v,len=%d,cap=%d\n", s1, len(s1), cap(s1))
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]=", arr[2:6]) //arr[2:6]= [2 3 4 5]
	fmt.Println("arr[:6]=", arr[:6])   //arr[:6]= [0 1 2 3 4 5]
	fmt.Println("arr[2:]=", arr[2:])   //arr[2:]= [2 3 4 5 6 7]
	fmt.Println("arr[:]=", arr[:])     //arr[:]= [0 1 2 3 4 5 6 7]
	s1 := arr[2:]
	s2 := arr[:]
	updateSlice(s1) //引用传递
	fmt.Println(s1) //[100 3 4 5 6 7]
	fmt.Println(s2) //[0 1 100 3 4 5 6 7]

	updateSlice(s2) //引用传递
	fmt.Println(s1) //[100 3 4 5 6 7]
	fmt.Println(s2) //[100 1 100 3 4 5 6 7]
	fmt.Println("slice扩展")
	s3 := arr[2:6] //[100 3 4 5]
	s4 := s3[3:5]  //[5 6]
	fmt.Printf("s3=%v,len(s3)=%d,cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4=%v,len(s4)=%d,cap(s4)=%d\n", s4, len(s4), cap(s4))
	//添加元素
	fmt.Println("添加元素")
	s3 = append(s3, s2...)
	s4 = append(s4, 10)
	fmt.Printf("s3=%v,len(s3)=%d,cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4=%v,len(s4)=%d,cap(s4)=%d\n", s4, len(s4), cap(s4))
	/*
		s3=[100 3 4 5 100 1 100 3 4 5 6 7],len(s3)=12,cap(s3)=12
		s4=[5 6 10],len(s4)=3,cap(s4)=3
	*/
	//slice操作
	var s []int //Zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	/*
		[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53 55 57 59 61 63 65 67 69 71 73 75 77 79 81 83 85 87 89 91 93 95 97 99 101 103 105 107 109 111 113 115 117 119 121 123 125 127 129 131 133 135 137 139 141 143 145 147 149 151 153 155 157 159 161 163 165 167 169 171 173 175 177 179 181 183 185 187 189 191 193 195 197 199]
	*/
	s5 := []int{2, 4, 6, 8}
	printSlice(s5) //s1=[2 4 6 8],len=4,cap=4
	s6 := make([]int, 16)

	s7 := make([]int, 10, 32)
	printSlice(s6)               //s1=[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0],len=16,cap=16
	printSlice(s7)               //s1=[0 0 0 0 0 0 0 0 0 0],len=10,cap=32
	fmt.Println("复制切片")
	copy(s6, s5)
	printSlice(s6) //s1=[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0],len=16,cap=16
	fmt.Println("删除切片元素")
	fmt.Println(s5[:1], s5[3:])
	s5 = append(s5[:1], s5[2:]...)
	fmt.Println(s5)//[2 6 8]删除4
	fmt.Println("Poping from front")
	front := s5[0] //2
	s5 = s5[1:] //[6 8]
	fmt.Println(s5)
	fmt.Println("Poping from back")
	tail := s5[len(s5)-1] //8
	s5 = s5[:len(s5)-1]//[6]
	fmt.Println(s5)//[6]
	fmt.Println(front,tail)
	printSlice(s5)//s1=[6],len=1,cap=3
}
